import logo from './logo.svg';
import styles from './App.module.css';
import { createSignal } from 'solid-js';

function App() {
  const [res, setRes] = createSignal('');
  const fetchAPI = () => {
    fetch('http://localhost:8080/api')
      .then((res) => res.json())
      .then((data) => setRes(data));
  };

  return (
    <div class={styles.App}>
      <header class={styles.header}>
        <img src={logo} class={styles.logo} alt='logo' />
        <p>
          Edit <code>src/App.jsx</code> and save to reload.
        </p>
        <a
          class={styles.link}
          href='https://github.com/solidjs/solid'
          target='_blank'
          rel='noopener noreferrer'
        >
          Learn Solid
        </a>
        <button onClick={() => fetchAPI()}>Fetch server</button>
        <p>{res()}</p>
      </header>
    </div>
  );
}

export default App;
