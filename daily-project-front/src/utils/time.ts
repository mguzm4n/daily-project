export const formatTweetDate = (dateString: string) => {
  const date = new Date(dateString);
  const dateTimeStr = date.toLocaleString([], { 
    hour: 'numeric', minute: 'numeric',  hour12: true,
    day: 'numeric', month: 'short', year: 'numeric'
  });

  const splitDatetime = dateTimeStr.split(",");
  const timeStr = splitDatetime[1].trim()
  const dateStr = splitDatetime[0]

  return `${timeStr} · ${dateStr}`;
};