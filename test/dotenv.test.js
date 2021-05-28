/**
 * MIT License
 *
 * Copyright (c) 2021 Iván Szkiba
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

import { group } from "k6";

export { options } from "./expect.js";
import { describe } from "./expect.js";
import dotenv from "k6/x/dotenv";

const sample = open("./testdata/sample.env");

const object = `
FOO = bar
SIZE = small
`;

export default function () {
  describe("parse", (t) => {
    const obj = dotenv.parse(object);
    t.expect(obj.FOO).as("FOO").toEqual("bar");
    t.expect(obj.SIZE).as("SIZE").toEqual("small");
  });

  describe("stringify", (t) => {
    const obj = { size: "small", foo: "bar" };
    const str = dotenv.stringify(obj);
    t.expect(str).as("text").toEqual(`foo="bar"\nsize="small"`);
  });

  describe("sample", (t) => {
    const obj = dotenv.parse(sample);
    t.expect(obj.FOO).as("FOO").toEqual("bar");
    t.expect(obj.SIZE).as("SIZE").toEqual("small");
  });
}
