//! blobstore implementation
//!

mod blobstore;
pub use blobstore::*;


impl BlobstoreResult {

    pub fn ok() -> Self {
        BlobstoreResult { success: true, error: None }
    }

    pub fn error<T: ToString>(s: T) -> Self {
        BlobstoreResult { success: false, error: Some(s.to_string()) }
    }
}