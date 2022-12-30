import { fizzBuzz } from "./src/fizzbuzz.ts";

if (import.meta.main) {
  for (let i = 0; i < 30; i++) {
    console.log(fizzBuzz(i));
  }
}
