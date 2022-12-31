import { fizzBuzz } from "../src/fizzbuzz.ts";
import { assertEquals } from "https://deno.land/std@0.149.0/testing/asserts.ts";

Deno.test({
  name: "15: fizzbuzz",
  fn: () => {
    assertEquals(fizzBuzz(15), "FIZZBUZZ");
  },
});

Deno.test({
  name: "3: fizz",
  fn: () => {
    assertEquals(fizzBuzz(15), "FIZZ");
  },
});
