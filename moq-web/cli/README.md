# `moq-web/cli`

Command-line utilities for `@qumo/moq`.

## Subdirectories

- [`interop/`](interop/): TypeScript interop client implementations for testing Media over QUIC (MOQ Lite) over WebTransport with Deno.
  - `client.ts`: Deno WebTransport client implementation.
  - `main.ts`: CLI entry point for executing interop tests.
  - `run_secure.ts`: Certificate hash calculation wrapper for self-signed WebTransport TLS testing.
