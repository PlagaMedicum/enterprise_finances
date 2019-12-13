function pad(num, len) {
  let s = String(num);
  while (s.length < len) s = "0" + s;
  return s
}

function timeConverter(day, month, year) {
  return `${pad(year, 4)}-${pad(month, 2)}-${pad(day, 2)}T00:00:00Z`
}

export default timeConverter;
