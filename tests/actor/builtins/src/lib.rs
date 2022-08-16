use wasmbus_rpc::actor::prelude::*;
use wasmcloud_interface_logging::{debug, error, info, warn};
use wasmcloud_interface_numbergen::{generate_guid, random_32, random_in_range};
use wasmcloud_test_util::{
    check, check_eq, run_selected,
    testing::{TestOptions, TestResult, Testing, TestingReceiver },
};

#[derive(Debug, Default, Actor, HealthResponder)]
#[services(Actor, Testing)]
struct NumbergenTestActor {}

#[async_trait]
impl Testing for NumbergenTestActor {
    async fn start(&self, _ctx: &Context, opts: &TestOptions) -> RpcResult<Vec<TestResult>> {
        console_log(&format!("numbergen test actor starting"));
        let results = run_selected!(
            opts,
            numbergen_uuid,
            numbergen_rand32,
            numbergen_rand_range,
            logging_levels,
        );
        Ok(results)
    }

    async fn foo(&self, _ctx: &Context) -> RpcResult<wasmcloud_test_util::testing::SampleUnion> { Err(RpcError::NotImplemented) }
}

/// test uuid generation
async fn numbergen_uuid(_opt: &TestOptions) -> RpcResult<()> {
    // generate twice and confirm they aren't the same
    // confirm that generated uuid is lowercase (per documentation)
    let uuid = generate_guid().await?;
    check_eq!(uuid.len(), 36)?;

    let uuid2 = generate_guid().await?;

    check!(uuid != uuid2)?;

    check_eq!(&uuid.to_lowercase(), &uuid)?;

    Ok(())
}

/// test random 32-bit generation
async fn numbergen_rand32(_opt: &TestOptions) -> RpcResult<()> {
    let mut min_val = u32::MAX;
    let mut max_val: u32 = 0;
    let mut sum: u64 = 0;
    const COUNT: usize = 24;
    const MARGIN: u32 = 14;

    for _ in 0..COUNT {
        let v = random_32().await?;
        min_val = std::cmp::min(min_val, v);
        max_val = std::cmp::max(max_val, v);
        sum += v as u64;
    }

    // cheap randomness tests:
    //   - not all zero
    //   - not all the same
    //   - mean value close to 2^31
    check!(sum > 0)?; // not all zero
    check!(min_val != max_val)?; // not all the same
    check!(min_val < u32::MAX)?;
    check!(max_val > 0)?;
    let avg = (sum / (COUNT as u64)) as u32;
    let spread = ((avg >> 26) as i32 - 32i32).abs() as u32; // should be close to zero
    if spread > MARGIN {
        return Err(RpcError::Other(format!("should be approx 0: {}", spread)));
    }

    Ok(())
}

async fn numbergen_rand_range(_opt: &TestOptions) -> RpcResult<()> {
    const MIN_VAL: u32 = 5;
    const MAX_VAL: u32 = 8;
    const TOTAL: usize = 20;
    let mut sum: u32 = 0;

    for _ in 0..TOTAL {
        let v = random_in_range(MIN_VAL, MAX_VAL).await?;

        check!(v >= MIN_VAL)?;
        check!(v <= MAX_VAL)?;
        sum += v;
    }
    check!(sum != MIN_VAL * (TOTAL as u32))?;
    check!(sum != MAX_VAL * (TOTAL as u32))?;
    Ok(())
}

async fn logging_levels(_opt: &TestOptions) -> RpcResult<()> {
    debug!("This message is at debug level");
    info!("This message is at info level");
    warn!("This message is at warn level");
    error!("This message is at error level");

    info!("This should say ABC: {}{}{}", "A", "B", "C");
    Ok(())
}
