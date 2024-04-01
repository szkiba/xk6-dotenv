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

## Usage

The `.env` files are loaded automatically when starting k6. To use it, simply create the appropriate `.env` file (see table above) and set (optional) the `K6_ENV` environment variable.

For the most convenient use, create a file called `.env.local` and write the environment variables you want to set in it. One variable per line, in `name=value` form.

```sh
SOME_ENV_VAR=somevalue
```

If you want to be really fancy with your env file you can do comments and exports:

```sh
# I am a comment and that is OK
SOME_VAR=someval
FOO=BAR # comments at line end are OK too
export BAR=BAZ
```

## Download

You can download pre-built k6 binaries from [Releases](https://github.com/szkiba/xk6-dotenv/releases/) page. Check [Packages](https://github.com/szkiba/xk6-dotenv/pkgs/container/xk6-dotenv) page for pre-built k6 Docker images.

## Build

The [xk6](https://github.com/grafana/xk6) build tool can be used to build a k6 that will include xk6-faker extension:

```bash
$ xk6 build --with github.com/szkiba/xk6-dotenv@latest
```

For more build options and how to use xk6, check out the [xk6 documentation](https://github.com/grafana/xk6).

