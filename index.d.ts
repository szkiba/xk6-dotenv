/**
 * xk6-dotenv loads env vars from a .env file.
 *
 * This extension follow the [convention](https://github.com/bkeepers/dotenv#what-other-env-files-can-i-use) for managing multiple environments (i.e. development, test, production). The actual environment name came from an env variable named `K6_ENV`. Setting this to `false` value disable the convention mentoined above and no dotenv file will be loaded automatically.
 *
 * | Hierarchy Priority | Filename                 | K6_ENV                 | Should I `.gitignore`it? | Notes                                                        |
 * | ------------------ | ------------------------ | ---------------------- | ------------------------ | ------------------------------------------------------------ |
 * | 1st (highest)      | `.env.development.local` | development            | Yes!                     | Local overrides of environment-specific settings.            |
 * | 1st                | `.env.test.local`        | test                   | Yes!                     | Local overrides of environment-specific settings.            |
 * | 1st                | `.env.production.local`  | production             | Yes!                     | Local overrides of environment-specific settings.            |
 * | 2nd                | `.env.local`             | (any _expect_ `false`) | Definitely.              | Local overrides. This file is loaded for all environments _except_ `test`. |
 * | 3rd                | `.env.development`       | development            | No.                      | Shared environment-specific settings                         |
 * | 3rd                | `.env.test`              | test                   | No.                      | Shared environment-specific settings                         |
 * | 3rd                | `.env.production`        | production             | No.                      | Shared environment-specific settings                         |
 * | Last               | `.env`                   | (any _expect_ `false`) | Depends                  | The Original                                                 |
 *
 * ## Usage
 *
 * Import an entire module's contents:
 * ```JavaScript
 * import * as dotenv from "k6/x/dotenv";
 * ```
 *
 * Import a single export from a module:
 * ```JavaScript
 * import { parse } from "k6/x/dotenv";
 * ```
 */

/**
 * The parse() method parses a .env format string, constructing the JavaScript object described by the string.
 *
 * @param text The string to parse as .env
 * @returns The Object corresponding to the given .env variables.
 */
export declare function parse(text: string): Record<string, string>;

/**
 * The stringify() method converts a JavaScript object to a .env format string.
 *
 * @param value The value to convert to a .env string
 * @returns A .env format string representing the given object
 */
export declare function stringify(value: Record<string, string>): string;
