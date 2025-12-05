# xk6-dotenv

A k6 extension that loads env vars from a .env file.

## Why Use .env Files?

Loading environment variables from `.env` files provides several key benefits for k6 testing:

- **Secure Secrets Management** - Keep sensitive data like API keys, passwords, and tokens out of your test scripts and version control. Store them locally in `.env.local` files that are gitignored.
- **Environment-Specific Configuration** - Easily switch between development, staging, and production environments without modifying your test scripts. Each environment can have its own configuration file.
- **Team Collaboration** - Share default configuration via committed `.env` files while allowing team members to override settings locally without conflicts.
- **Simplified Test Execution** - No need to pass multiple `-e` flags to k6 or manually export environment variables before running tests.

## How It Works

This extension follows the [convention](https://github.com/bkeepers/dotenv#what-other-env-files-can-i-use) for managing multiple environments (i.e. development, test, production). The actual environment name comes from an env variable named `K6_ENV`. Setting this to `false` value disables the convention mentioned above and no .env file will be loaded automatically.

> [!NOTE]
> This extension doesn't add any JavaScript API to k6. It simply extends k6's behavior by automatically loading environment variables from `.env` files at startup. Your k6 test scripts can access these variables in the usual way using the `__ENV` map.

| Hierarchy Priority | Filename                 | K6_ENV                 | Should I `.gitignore`it? | Notes                                                        |
| ------------------ | ------------------------ | ---------------------- | ------------------------ | ------------------------------------------------------------ |
| 1st (highest)      | `.env.development.local` | development            | Yes!                     | Local overrides of environment-specific settings.            |
| 1st                | `.env.test.local`        | test                   | Yes!                     | Local overrides of environment-specific settings.            |
| 1st                | `.env.production.local`  | production             | Yes!                     | Local overrides of environment-specific settings.            |
| 2nd                | `.env.local`             | (any _except_ `false`) | Definitely.              | Local overrides. This file is loaded for all environments _except_ `test`. |
| 3rd                | `.env.development`       | development            | No.                      | Shared environment-specific settings                         |
| 3rd                | `.env.test`              | test                   | No.                      | Shared environment-specific settings                         |
| 3rd                | `.env.production`        | production             | No.                      | Shared environment-specific settings                         |
| Last               | `.env`                   | (any _except_ `false`) | Depends                  | The Original                                                 |

## Usage

The `.env` files are loaded automatically when starting k6. To use it, simply create the appropriate `.env` file (see table above) and set (optional) the `K6_ENV` environment variable.

For the most convenient use, create a file called `.env.local` and write the environment variables you want to set in it. One variable per line, in `name=value` form.

```sh
API_KEY=your-api-key-here
BASE_URL=https://example.com
```

Then access the variables in your k6 test script using the `__ENV` map:

```javascript
import http from 'k6/http';
import { check } from 'k6';

export default function () {
  const apiKey = __ENV.API_KEY;
  const baseUrl = __ENV.BASE_URL;
  
  const res = http.get(`${baseUrl}/api/data`, {
    headers: { 'Authorization': `Bearer ${apiKey}` }
  });
  
  check(res, { 'status is 200': (r) => r.status === 200 });
}
```

### .env File Format

Your `.env` files support comments and export statements:

```sh
# API Configuration
API_KEY=somevalue
BASE_URL=https://test.example.com # inline comments are supported
export TIMEOUT=30
```

## Download

You can download pre-built k6 binaries from [Releases](https://github.com/szkiba/xk6-dotenv/releases/) page. Check [Packages](https://github.com/szkiba/xk6-dotenv/pkgs/container/xk6-dotenv) page for pre-built k6 Docker images.

## Build

The [xk6](https://github.com/grafana/xk6) build tool can be used to build a k6 that will include xk6-dotenv extension:

```bash
$ xk6 build --with github.com/szkiba/xk6-dotenv@latest
```

For more build options and how to use xk6, check out the [xk6 documentation](https://github.com/grafana/xk6).

