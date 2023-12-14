/// If this is to be executed in a blockchain context, then we need to delegate these hashing
/// functions to a native implementation through host function calls.
/// This trait provides that interface.
pub trait HostFunctionsProvider {
    /// The SHA-256 hash algorithm
    fn sha2_256(message: &[u8]) -> [u8; 32];

    /// The SHA-512 hash algorithm
    fn sha2_512(message: &[u8]) -> [u8; 64];

    /// The SHA-512 hash algorithm with its output truncated to 256 bits.
    fn sha2_512_truncated(message: &[u8]) -> [u8; 32];

    /// The Keccak-256 hash function.
    fn keccak_256(message: &[u8]) -> [u8; 32];

    /// The Ripemd160 hash function.
    fn ripemd160(message: &[u8]) -> [u8; 20];

    /// BLAKE2b-512 hash function.
    fn blake2b_512(message: &[u8]) -> [u8; 64];

    /// BLAKE2s-256 hash function.
    fn blake2s_256(message: &[u8]) -> [u8; 32];

    /// BLAKE3 hash function.
    fn blake3(message: &[u8]) -> [u8; 32];
}

#[cfg(any(feature = "host-functions", test))]
pub mod host_functions_impl {
    use crate::host_functions::HostFunctionsProvider;
    use blake2::{Blake2b512, Blake2s256};
    use ripemd::Ripemd160;
    use sha2::{Digest, Sha256, Sha512, Sha512_256};
    use sha3::Keccak256;

    pub struct HostFunctionsManager;
    impl HostFunctionsProvider for HostFunctionsManager {
        fn sha2_256(message: &[u8]) -> [u8; 32] {
            let digest = Sha256::digest(message);
            let mut buf = [0u8; 32];
            buf.copy_from_slice(&digest);
            buf
        }

        fn sha2_512(message: &[u8]) -> [u8; 64] {
            let digest = Sha512::digest(message);
            let mut buf = [0u8; 64];
            buf.copy_from_slice(&digest);
            buf
        }

        fn sha2_512_truncated(message: &[u8]) -> [u8; 32] {
            let digest = Sha512_256::digest(message);
            let mut buf = [0u8; 32];
            buf.copy_from_slice(&digest);
            buf
        }

        fn keccak_256(message: &[u8]) -> [u8; 32] {
            let digest = Keccak256::digest(message);
            let mut buf = [0u8; 32];
            buf.copy_from_slice(&digest);
            buf
        }

        fn ripemd160(message: &[u8]) -> [u8; 20] {
            let digest = Ripemd160::digest(message);
            let mut buf = [0u8; 20];
            buf.copy_from_slice(&digest);
            buf
        }

        fn blake2b_512(message: &[u8]) -> [u8; 64] {
            let digest = Blake2b512::digest(message);
            let mut buf = [0u8; 64];
            buf.copy_from_slice(&digest);
            buf
        }

        fn blake2s_256(message: &[u8]) -> [u8; 32] {
            let digest = Blake2s256::digest(message);
            let mut buf = [0u8; 32];
            buf.copy_from_slice(&digest);
            buf
        }

        fn blake3(message: &[u8]) -> [u8; 32] {
            blake3::hash(message).into()
        }
    }
}
