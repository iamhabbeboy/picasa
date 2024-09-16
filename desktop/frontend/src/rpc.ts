import { EventsOnMultiple } from "../wailsjs/runtime";

function on(eventName: string, callback: (...data: any) => void) {
  EventsOnMultiple(eventName, callback, -1);
}

const rpc = { on };
export default rpc;
