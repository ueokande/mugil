export function fetch(date) {
  return {
    type: "TASKS_FETCH",
    entries: []
  };
}

export function update(entries) {
  return {
    type: "TASKS_UPDATE",
    entries: entries
  };
}
