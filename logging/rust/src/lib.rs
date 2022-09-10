//! logging capability provider
//!
//! ```no_compile
//!   use wasmcloud_interface_logging::{log,debug,info};
//!   log("info", "log a message at info level");
//!   info!("this is also at {} level", "info");
//! ```

mod logging;
pub use logging::*;

#[cfg(target_arch = "wasm32")]
pub use wasm_log::log;

#[cfg(target_arch = "wasm32")]
mod wasm_log {
    use std::string::ToString;
    use wasmbus_rpc::error::RpcResult;

    /// log a text message at a severity level (debug, info, warn, or error)
    /// parameters may be `&str`, `String`, or `&String`.
    /// You can also use this crate's macros `debug`,`info`,`warn`, and `error`
    pub async fn log<T1: ToString, T2: ToString>(level: T1, text: T2) -> RpcResult<()> {
        use crate::logging::{Logging, LoggingSender};
        let entry = crate::logging::LogEntry {
            level: level.to_string(),
            text: text.to_string(),
        };
        let ctx = wasmbus_rpc::common::Context::default();
        let logger = LoggingSender::new();
        logger.write_log(&ctx, &entry).await
    }

    #[macro_export]
    #[doc(hidden)]
    /// internal macro to convert format arts into a string
    macro_rules! log_fmt {
        ( $($arg:expr),* $(,)? ) => ({ format!("{}", format_args!($($arg),*) ) });
    }

    /// Emit log at 'debug' level. The syntax for this macro is the same as `log::debug`.
    /// Note that parameters are always evaluated at any logging level
    #[cfg(not(feature = "sync_macro"))]
    #[macro_export]
    macro_rules! debug {
        ( $($arg:expr),* $(,)? ) => ({ let _ = $crate::log("debug", $crate::log_fmt!($($arg),*) ).await?; });
    }

    /// Emit log at 'info' level. The syntax for this macro is the same as `log::info`.
    /// Note that parameters are always evaluated at any logging level
    #[cfg(not(feature = "sync_macro"))]
    #[macro_export]
    macro_rules! info {
        ( $($arg:expr),* $(,)? ) => ({ let _ = $crate::log("info", $crate::log_fmt!($($arg),*) ).await?; });
    }

    /// Emit log at 'warn' level. The syntax for this macro is the same as `log::warn`.
    /// Note that parameters are always evaluated at any logging level
    #[cfg(not(feature = "sync_macro"))]
    #[macro_export]
    macro_rules! warn {
        ( $($arg:expr),* $(,)? ) => ({ let _ = $crate::log("warn", $crate::log_fmt!($($arg),*) ).await?; });
    }

    /// Emit log at 'error' level. The syntax for this macro is the same as `log::error`.
    /// 'error' is the highest priority level.
    #[cfg(not(feature = "sync_macro"))]
    #[macro_export]
    macro_rules! error {
        ( $($arg:expr),* $(,)? ) => ({ let _ = $crate::log("error", $crate::log_fmt!($($arg),*) ).await?; });
    }
}

#[cfg(all(feature = "sync_macro", target_arch = "wasm32"))]
pub mod sync {

    use serde::Serialize;
    use wasmbus_rpc::{actor::prelude::host_call, common::serialize};

    #[derive(Serialize)]
    struct LogEntryB<'a> {
        level: &'a str,
        text: &'a str,
    }

    pub fn log(level: &str, text: &str) {
        let level = if ["debug", "info", "warn", "error"].contains(&level) {
            level
        } else {
            "info"
        };
        serialize(&LogEntryB { level, text })
            .and_then(|data| {
                host_call(
                    "default",
                    "wasmcloud:builtin:logging",
                    "Logging.WriteLog",
                    &data,
                )
            })
            .ok();
    }

    /// Emit log at 'debug' level. The syntax for this macro is the same as `log::debug`.
    /// Note that parameters are always evaluated at any logging level
    #[macro_export]
    macro_rules! debug { ( $($arg:expr),* $(,)? ) => ({ let _ = $crate::sync::log("debug", &$crate::log_fmt!($($arg),*) ); }); }

    /// Emit log at 'info' level. The syntax for this macro is the same as `log::info`.
    /// Note that parameters are always evaluated at any logging level
    #[macro_export]
    macro_rules! info { ( $($arg:expr),* $(,)? ) => ({ let _ = $crate::sync::log("info", &$crate::log_fmt!($($arg),*) ); }); }

    /// Emit log at 'warn' level. The syntax for this macro is the same as `log::warn`.
    /// Note that parameters are always evaluated at any logging level
    #[macro_export]
    macro_rules! warn { ( $($arg:expr),* $(,)? ) => ({ let _ = $crate::sync::log("warn", &$crate::log_fmt!($($arg),*) ); }); }

    /// Emit log at 'error' level. The syntax for this macro is the same as `log::error`.
    /// Note that parameters are always evaluated at any logging level
    #[macro_export]
    macro_rules! error { ( $($arg:expr),* $(,)? ) => ({ let _ = $crate::sync::log("error", &$crate::log_fmt!($($arg),*) ); }); }
}
