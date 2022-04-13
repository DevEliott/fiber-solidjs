import logo from './logo.svg';
import styles from './App.module.css';
import { createSignal } from 'solid-js';

function App() {
  const [res, setRes] = createSignal('');
  const [input, setInput] = createSignal('');
  const [id, setId] = createSignal('');
  const fetchAPI = () => {
    fetch('http://localhost:8080/api')
      .then((res) => res.json())
      .then((data) => setRes(data));
  };
  const ws = new WebSocket('ws://localhost:8080/ws');
  ws.onopen = (ev) => {
    console.log('WS connection opened');
  };
  ws.onmessage = (ev) => {
    const payload = JSON.parse(ev.data);
    console.log(payload);
    if (payload.action === 'ID') {
      setId(payload.data);
    } else setRes(payload);
  };

  const sendWS = (e) => {
    if (e.key === 'Enter') {
      ws.send(
        JSON.stringify({ data: input(), action: 'Message', senderId: id() }),
      );
      setInput('');
    }
  };

  return (
    <div class={styles.App}>
      <header class={styles.header}>
        <img src={logo} class={styles.logo} alt='logo' />
        <p>
          Edit <code>src/App.jsx</code> and save to reload.
        </p>
        <input
          type='text'
          value={input()}
          onInput={(e) => setInput(e.currentTarget.value)}
          onKeyDown={sendWS}
        />
        <button onClick={() => fetchAPI()}>Fetch server</button>
        <p>{res()}</p>
      </header>
    </div>
  );
}

export default App;
