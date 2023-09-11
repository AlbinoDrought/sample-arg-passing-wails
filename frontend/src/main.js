import './style.css';
import './app.css';

import logo from './assets/images/logo-universal.png';
import { EventsOn } from '../wailsjs/runtime/runtime';

document.querySelector('#app').innerHTML = `
    <img id="logo" class="logo">
      <div class="result" id="result">Run the app with some args, they will appear below 👇</div>
      <pre class="input-box" id="args"></pre>
    </div>
`;
document.getElementById('logo').src = logo;

const elArgs = document.getElementById('args');
EventsOn('instance-message', (args) => {
    elArgs.innerText += args + '\n';
});
