export function fizzBuzz(x: number): string {
  if (x % 15 === 0) {
    return "FIZZBUZZ";
  } else if (x % 5 === 0) {
    return "BUZZ";
  } else if (x % 3 === 0) {
    return "FIZZ";
  } else {
    return x.toString();
  }
}
