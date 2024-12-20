import { EventsEmit, EventsOnMultiple } from "../wailsjs/runtime";

function on(eventName: string, callback: (...data: any) => void) {
  EventsOnMultiple(eventName, callback, -1);
}

function emit(eventName: string, payload?: Record<string, any>) {
  EventsEmit(eventName, payload);
}

const rpc = { on, emit };

export default rpc;
