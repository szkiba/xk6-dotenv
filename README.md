# xk6-dotenv

A k6 extension that loads env vars from a .env file.

This extension follow the [convention](https://github.com/bkeepers/dotenv#what-other-env-files-can-i-use) for managing multiple environments (i.e. development, test, production). The actual environment name came from an env variable named `K6_ENV`. Setting this to `false` value disable the convention mentoined above and no .env file will be loaded automatically.

| Hierarchy Priority | Filename                 | K6_ENV                 | Should I `.gitignore`it? | Notes                                                        |
| ------------------ | ------------------------ | ---------------------- | ------------------------ | ------------------------------------------------------------ |
| 1st (highest)      | `.env.development.local` | development            | Yes!                     | Local overrides of environment-specific settings.            |
| 1st                | `.env.test.local`        | test                   | Yes!                     | Local overrides of environment-specific settings.            |
| 1st                | `.env.production.local`  | production             | Yes!                     | Local overrides of environment-specific settings.            |
| 2nd                | `.env.local`             | (any _expect_ `false`) | Definitely.              | Local overrides. This file is loaded for all environments _except_ `test`. |
| 3rd                | `.env.development`       | development            | No.                      | Shared environment-specific settings                         |
| 3rd                | `.env.test`              | test                   | No.                      | Shared environment-specific settings                         |
| 3rd                | `.env.production`        | production             | No.                      | Shared environment-specific settings                         |
| Last               | `.env`                   | (any _expect_ `false`) | Depends                  | The Original                                                 |


The underlying implementation is https://github.com/joho/godotenv

Built for [k6](https://go.k6.io/k6) using [xk6](https://github.com/grafana/xk6).

## Usage

Import an entire module's contents:
```JavaScript
import * as dotenv from "k6/x/dotenv";
```

Import a single export from a module:
```JavaScript
import { parse } from "k6/x/dotenv";
```

## API

This extension can be used as a library:

- [parse](docs/README.md#parse)
- [stringify](docs/README.md#stringify)

For complete API documentation click [here](docs/README.md)!

## Build

To build a `k6` binary with this extension, first ensure you have the prerequisites:

- [Go toolchain](https://go101.org/article/go-toolchain.html)
- Git

Then:

1. Install `xk6`:
  ```bash
  $ go install go.k6.io/xk6/cmd/xk6@latest
  ```

2. Build the binary:
  ```bash
  $ xk6 build --with github.com/szkiba/xk6-dotenv@latest
  ```
