const eventsResult = [];
const eventFn = (type, e) => {
  //   console.log("~~~ eventFn(): type: ", type);
  //   console.log("~~~ eventFn(): data: ", e?.data);
  eventsResult.push({ type, data: JSON.parse(e?.data) });
};
const eventsList = [
  "head",
  "finalized_checkpoint",
  "chain_reorg",
  "block",
  "attestation",
  "voluntary_exit",
  "contribution_and_proof",
];
let url = "http://localhost:5051/eth/v1/events?topics=";
eventsList.map((eventsItem) => {
  url += `${eventsItem},`;
});
url = url.substring(0, url.length - 1);
// console.log(url);
const eventSource = new EventSource(url);

eventsList.map((eventsItem) => {
  try {
    eventSource.addEventListener(eventsItem, (e) => eventFn(eventsItem, e));
  } catch (error) {
    console.error(`~~~ error: ${error}`);
  }
});
const getMatching = () => {
  const blockCount = {};
  eventsResult.map((eventsResultItem) => {
    const block = eventsResultItem?.data?.epoch;
    if (block) {
      if (blockCount[block] === undefined) {
        blockCount[block] = 0;
      } else {
        blockCount[block] = ++blockCount[block];
      }
    }
  });
  return blockCount;
};
//  getMatching()