//! mlinference Interface

mod mlinference;
pub use mlinference::*;

// flag bit 0 LSB: row-major order (0) column major order (1)
pub const TENSOR_FLAG_ROW_MAJOR: u8 = 0;
pub const TENSOR_FLAG_COLUMN_MAJOR: u8 = 1;
// flag bit 1: little-endian (0) or big endian (2)
pub const TENSOR_FLAG_LITTLE_ENDIAN: u8 = 0;
pub const TENSOR_FLAG_BIG_ENDIAN: u8 = 2;

impl Default for Tensor {
    fn default() -> Tensor {
        Tensor {
            value_types: vec![ValueType::ValueF32],
            dimensions: vec![],
            data: vec![],
            flags: TENSOR_FLAG_ROW_MAJOR | TENSOR_FLAG_LITTLE_ENDIAN,
        }
    }
}

impl Tensor {
    /// Returns dimensions as an array of usize
    pub fn shape(&self) -> Vec<usize> {
        self.dimensions.iter().map(|d| *d as usize).collect()
    }

    pub fn set_row_major(&mut self) {
        self.flags = (self.flags & 0xfe) | TENSOR_FLAG_ROW_MAJOR;
    }

    pub fn is_row_major(&self) -> bool {
        (self.flags & 1) == TENSOR_FLAG_ROW_MAJOR
    }

    pub fn set_column_major(&mut self) {
        self.flags = (self.flags & 0xfe) | TENSOR_FLAG_COLUMN_MAJOR;
    }

    pub fn is_column_major(&self) -> bool {
        (self.flags & 1) == TENSOR_FLAG_COLUMN_MAJOR
    }

    pub fn set_little_endian(&mut self) {
        self.flags = (self.flags & 0xfd) | TENSOR_FLAG_LITTLE_ENDIAN;
    }

    pub fn is_little_endian(&self) -> bool {
        (self.flags & 2) == TENSOR_FLAG_LITTLE_ENDIAN
    }

    pub fn set_big_endian(&mut self) {
        self.flags = (self.flags & 0xfd) | TENSOR_FLAG_BIG_ENDIAN;
    }

    pub fn is_big_endian(&self) -> bool {
        (self.flags & 2) == TENSOR_FLAG_BIG_ENDIAN
    }

    /// perform validation checking on dimensions
    pub fn check_dims(&self) -> Result<(), String> {
        if self.data.len() as u64 > (i32::MAX as u64) {
            return Err("data too large".into());
        }
        if self.data.is_empty() {
            return Err("data is empty".into());
        }
        let ndims = self.dimensions.len();
        let nvals = self.value_types.len();
        if ndims == 0 {
            return Err("'dimensions' is empty".into());
        }
        if nvals == 0 {
            return Err("'value_types' is empty".into());
        }
        if nvals != 1 && nvals != ndims {
            return Err(format!(
                "'value_types' should have length 1 or {} (# dimensions), not {}",
                ndims, nvals
            ));
        }
        if self.dimensions.iter().any(|d| *d == 0) {
            return Err("dimension length cannot be zero".into());
        }
        let mut bytes_total = 1usize;
        if nvals == 1 {
            for d in self.dimensions.iter() {
                bytes_total *= *d as usize;
            }
            bytes_total *= self.value_types.get(0).unwrap().datum_size();
        } else {
            for (d, v) in std::iter::zip(self.dimensions.iter(), self.value_types.iter()) {
                let dim_datum_size = v.datum_size();
                bytes_total *= *d as usize * dim_datum_size;
            }
        }
        if bytes_total != self.data.len() {
            let dim_sizes = self
                .value_types
                .iter()
                .map(|v| v.datum_size())
                .collect::<Vec<usize>>();
            return Err(format!(
                "Raw data size ({} bytes) does not fit dimensions ({:?}) * datum size ({:?})",
                self.data.len(),
                &self.dimensions,
                dim_sizes
            ));
        }
        Ok(())
    }
}

impl Default for Status {
    fn default() -> Status {
        Status::Success
    }
}

impl TryFrom<&str> for ValueType {
    type Error = String;

    fn try_from(value: &str) -> Result<Self, Self::Error> {
        match value {
            "u8" | "U8" => Ok(ValueType::ValueU8),
            "u16" | "U16" => Ok(ValueType::ValueU16),
            "u32" | "U32" => Ok(ValueType::ValueU32),
            "u64" | "U64" => Ok(ValueType::ValueU64),
            "u128" | "U128" => Ok(ValueType::ValueU128),

            "s8" | "S8" => Ok(ValueType::ValueS8),
            "s16" | "S16" => Ok(ValueType::ValueS16),
            "s32" | "S32" => Ok(ValueType::ValueS32),
            "s64" | "S64" => Ok(ValueType::ValueS64),
            "s128" | "S128" => Ok(ValueType::ValueS128),

            "f16" | "F16" => Ok(ValueType::ValueF16),
            "f32" | "F32" => Ok(ValueType::ValueF32),
            "f64" | "F64" => Ok(ValueType::ValueF64),
            "f128" | "F128" => Ok(ValueType::ValueF128),

            _ => Err(format!("invalid ValueType '{}'", value)),
        }
    }
}

impl ValueType {
    /// Return number of bytes per datum
    pub fn datum_size(&self) -> usize {
        match self {
            ValueType::ValueU8 | ValueType::ValueS8 => 1,

            ValueType::ValueU16 | ValueType::ValueS16 | ValueType::ValueF16 => 2,

            ValueType::ValueU32 | ValueType::ValueS32 | ValueType::ValueF32 => 4,

            ValueType::ValueU64 | ValueType::ValueS64 | ValueType::ValueF64 => 8,

            ValueType::ValueU128 | ValueType::ValueS128 | ValueType::ValueF128 => 16,
        }
    }
}
