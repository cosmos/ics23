extern crate failure;

pub type Result<T> = ::std::result::Result<T, failure::Error>;
pub type Hash = ::std::vec::Vec<u8>;
