const n = 5;

for (let i = 1; i <= n; i++) {
  let line = "";

  //space
  for (let j = 1; j <= n - i; j++) {
    line += " ";
  }

  //star
  for (let k = 1; k <= 2 * i - 1; k++) {
    line += "*";
  }

  console.log(line);
}
