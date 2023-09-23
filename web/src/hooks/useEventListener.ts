import { useEffect, useRef } from 'react';

type AllEventMaps = HTMLElementEventMap & DocumentEventMap & WindowEventMap;

const useEventListener = <K extends keyof AllEventMaps>(
  eventName: K,
  listener: (event: AllEventMaps[K]) => void,
  element?: HTMLElement | Document | Window,
  options?: boolean | AddEventListenerOptions,
) => {
  const listenerRef = useRef(listener);

  useEffect(() => {
    listenerRef.current = listener;
  }, [listener]);

  useEffect(() => {
    const el = element === undefined ? window : element;

    const eventListener = (event: AllEventMaps[K]) => listenerRef.current(event);

    el.addEventListener(eventName, eventListener as EventListenerOrEventListenerObject, options);

    return () => {
      el.removeEventListener(
        eventName,
        eventListener as EventListenerOrEventListenerObject,
        options,
      );
    };
  }, [eventName, element, options]);
};

export { useEventListener };
