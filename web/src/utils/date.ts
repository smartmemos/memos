// Helper function to convert proto2 Timestamp to Date
export const fromTimestamp = (t: { seconds: bigint; nanos: number }): Date => {
  let millis = Number(t.seconds || 0) * 1000;
  millis += (t.nanos || 0) / 1000000;
  return new Date(millis);
};