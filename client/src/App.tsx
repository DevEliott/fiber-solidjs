import { Component, createSignal } from 'solid-js';

interface Packet {
  action: string;
  data?: any;
  senderId?: string;
}

const App: Component = () => {
  const [res, setRes] = createSignal<Packet>();
  const [input, setInput] = createSignal('');
  const [id, setId] = createSignal('');
  const ws = new WebSocket('ws://localhost:8080/ws');
  ws.onopen = (ev) => {
    console.log('WS connection opened');
  };
  ws.onmessage = (ev) => {
    const payload: Packet = JSON.parse(ev.data);
    console.log(payload);
    if (payload.action === 'ID') {
      setId(payload.data);
    } else setRes(payload);
  };

  const sendWS = (e: KeyboardEvent) => {
    if (e.key === 'Enter') {
      ws.send(
        JSON.stringify({ data: input(), action: 'Message', senderId: id() }),
      );
      setInput('');
    }
  };

  return (
    <div class='min-h-screen bg-dark-200 grid place-items-center'>
      <div>
        <input
          placeholder='Enter your name'
          type='text'
          value={input()}
          onInput={(e) => setInput(e.currentTarget.value)}
          onKeyPress={(e) => sendWS(e)}
          class='rounded-lg font-medium p-6 text-5xl'
        />
        <pre>{JSON.stringify(res())}</pre>
      </div>
    </div>
  );
};

export default App;
