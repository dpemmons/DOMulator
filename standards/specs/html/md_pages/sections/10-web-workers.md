## [10]{.secno} Web workers[](#workers){.self-link} {#workers dfn-type="dfn" lt="web worker" export=""}

:::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Web_Workers_API](https://developer.mozilla.org/en-US/docs/Web/API/Web_Workers_API "Web Workers makes it possible to run a script operation in a background thread separate from the main execution thread of a web application. The advantage of this is that laborious processing can be performed in a separate thread, allowing the main (usually the UI) thread to run without being blocked/slowed down.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome2+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::

::: feature
[Web_Workers_API/Using_web_workers](https://developer.mozilla.org/en-US/docs/Web/API/Web_Workers_API/Using_web_workers "Web Workers are a simple means for web content to run scripts in background threads. The worker thread can perform tasks without interfering with the user interface. In addition, they can perform I/O using XMLHttpRequest (although the responseXML and channel attributes are always null) or fetch (with no such restrictions). Once created, a worker can send messages to the JavaScript code that created it by posting messages to an event handler specified by that code (and vice versa).")
:::
::::::

### [10.1]{.secno} Introduction[](#introduction-14){.self-link} {#introduction-14}

#### [10.1.1]{.secno} Scope[](#scope-2){.self-link} {#scope-2}

*This section is non-normative.*

This specification defines an API for running scripts in the background
independently of any user interface scripts.

This allows for long-running scripts that are not interrupted by scripts
that respond to clicks or other user interactions, and allows long tasks
to be executed without yielding to keep the page responsive.

Workers (as these background scripts are called herein) are relatively
heavy-weight, and are not intended to be used in large numbers. For
example, it would be inappropriate to launch one worker for each pixel
of a four megapixel image. The examples below show some appropriate uses
of workers.

Generally, workers are expected to be long-lived, have a high start-up
performance cost, and a high per-instance memory cost.

#### [10.1.2]{.secno} Examples[](#examples-6){.self-link} {#examples-6}

*This section is non-normative.*

There are a variety of uses that workers can be put to. The following
subsections show various examples of this use.

##### [10.1.2.1]{.secno} A background number-crunching worker[](#a-background-number-crunching-worker){.self-link}

*This section is non-normative.*

The simplest use of workers is for performing a computationally
expensive task without interrupting the user interface.

In this example, the main document spawns a worker to (naïvely) compute
prime numbers, and progressively displays the most recently found prime
number.

The main page is as follows:

``` html
<!DOCTYPE HTML>
<html lang="en">
 <head>
  <meta charset="utf-8">
  <title>Worker example: One-core computation</title>
 </head>
 <body>
  <p>The highest prime number discovered so far is: <output id="result"></output></p>
  <script>
   var worker = new Worker('worker.js');
   worker.onmessage = function (event) {
     document.getElementById('result').textContent = event.data;
   };
  </script>
 </body>
</html>
```

The
[`Worker()`{#a-background-number-crunching-worker:dom-worker}](#dom-worker)
constructor call creates a worker and returns a
[`Worker`{#a-background-number-crunching-worker:worker}](#worker) object
representing that worker, which is used to communicate with the worker.
That object\'s
[`onmessage`{#a-background-number-crunching-worker:handler-messageeventtarget-onmessage}](web-messaging.html#handler-messageeventtarget-onmessage)
event handler allows the code to receive messages from the worker.

The worker itself is as follows:

``` js
var n = 1;
search: while (true) {
  n += 1;
  for (var i = 2; i <= Math.sqrt(n); i += 1)
    if (n % i == 0)
     continue search;
  // found a prime!
  postMessage(n);
}
```

The bulk of this code is simply an unoptimized search for a prime
number. The
[`postMessage()`{#a-background-number-crunching-worker:dom-dedicatedworkerglobalscope-postmessage}](#dom-dedicatedworkerglobalscope-postmessage)
method is used to send a message back to the page when a prime is found.

[View this example online](/demos/workers/primes/page.html).

##### [10.1.2.2]{.secno} Using a JavaScript module as a worker[](#module-worker-example){.self-link} {#module-worker-example}

*This section is non-normative.*

All of our examples so far show workers that run [classic
scripts](webappapis.html#classic-script){#module-worker-example:classic-script}.
Workers can instead be instantiated using [module
scripts](webappapis.html#module-script){#module-worker-example:module-script},
which have the usual benefits: the ability to use the JavaScript
`import` statement to import other modules; strict mode by default; and
top-level declarations not polluting the worker\'s global scope.

As the `import` statement is available, the
[`importScripts()`{#module-worker-example:dom-workerglobalscope-importscripts}](#dom-workerglobalscope-importscripts)
method will automatically fail inside module workers.

In this example, the main document uses a worker to do off-main-thread
image manipulation. It imports the filters used from another module.

The main page is as follows:

``` html
<!DOCTYPE html>
<html lang="en">
<meta charset="utf-8">
<title>Worker example: image decoding</title>

<p>
  <label>
    Type an image URL to decode
    <input type="url" id="image-url" list="image-list">
    <datalist id="image-list">
      <option value="https://html.spec.whatwg.org/images/drawImage.png">
      <option value="https://html.spec.whatwg.org/images/robots.jpeg">
      <option value="https://html.spec.whatwg.org/images/arcTo2.png">
    </datalist>
  </label>
</p>

<p>
  <label>
    Choose a filter to apply
    <select id="filter">
      <option value="none">none</option>
      <option value="grayscale">grayscale</option>
      <option value="brighten">brighten by 20%</option>
    </select>
  </label>
</p>

<div id="output"></div>

<script type="module">
  const worker = new Worker("worker.js", { type: "module" });
  worker.onmessage = receiveFromWorker;

  const url = document.querySelector("#image-url");
  const filter = document.querySelector("#filter");
  const output = document.querySelector("#output");

  url.oninput = updateImage;
  filter.oninput = sendToWorker;

  let imageData, context;

  function updateImage() {
    const img = new Image();
    img.src = url.value;

    img.onload = () => {
      const canvas = document.createElement("canvas");
      canvas.width = img.width;
      canvas.height = img.height;

      context = canvas.getContext("2d");
      context.drawImage(img, 0, 0);
      imageData = context.getImageData(0, 0, canvas.width, canvas.height);

      sendToWorker();
      output.replaceChildren(canvas);
    };
  }

  function sendToWorker() {
    worker.postMessage({ imageData, filter: filter.value });
  }

  function receiveFromWorker(e) {
    context.putImageData(e.data, 0, 0);
  }
</script>
```

The worker file is then:

``` js
import * as filters from "./filters.js";

self.onmessage = e => {
  const { imageData, filter } = e.data;
  filters[filter](imageData);
  self.postMessage(imageData, [imageData.data.buffer]);
};
```

Which imports the file `filters.js`:

``` js
export function none() {}

export function grayscale({ data: d }) {
  for (let i = 0; i < d.length; i += 4) {
    const [r, g, b] = [d[i], d[i + 1], d[i + 2]];

    // CIE luminance for the RGB
    // The human eye is bad at seeing red and blue, so we de-emphasize them.
    d[i] = d[i + 1] = d[i + 2] = 0.2126 * r + 0.7152 * g + 0.0722 * b;
  }
};

export function brighten({ data: d }) {
  for (let i = 0; i < d.length; ++i) {
    d[i] *= 1.2;
  }
};
```

[View this example online](/demos/workers/modules/page.html).

##### [10.1.2.3]{.secno} Shared workers introduction[](#shared-workers-introduction){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[SharedWorker](https://developer.mozilla.org/en-US/docs/Web/API/SharedWorker "The SharedWorker interface represents a specific kind of worker that can be accessed from several browsing contexts, such as several windows, iframes or even workers. They implement an interface different than dedicated workers and have a different global scope, SharedWorkerGlobalScope.")

Support in all current engines.

::: support
[Firefox29+]{.firefox .yes}[Safari16+]{.safari .yes}[Chrome5+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android33+]{.firefox_android .yes}[Safari iOS16+]{.safari_ios
.yes}[Chrome AndroidNo]{.chrome_android .no}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet4.0--5.0]{.samsunginternet_android .no}[Opera
Android11--14]{.opera_android .no}
:::
::::
:::::

*This section is non-normative.*

This section introduces shared workers using a Hello World example.
Shared workers use slightly different APIs, since each worker can have
multiple connections.

This first example shows how you connect to a worker and how a worker
can send a message back to the page when it connects to it. Received
messages are displayed in a log.

Here is the HTML page:

``` html
<!DOCTYPE HTML>
<html lang="en">
<meta charset="utf-8">
<title>Shared workers: demo 1</title>
<pre id="log">Log:</pre>
<script>
  var worker = new SharedWorker('test.js');
  var log = document.getElementById('log');
  worker.port.onmessage = function(e) { // note: not worker.onmessage!
    log.textContent += '\n' + e.data;
  }
</script>
```

Here is the JavaScript worker:

``` js
onconnect = function(e) {
  var port = e.ports[0];
  port.postMessage('Hello World!');
}
```

[View this example online](/demos/workers/shared/001/test.html).

------------------------------------------------------------------------

This second example extends the first one by changing two things: first,
messages are received using `addEventListener()` instead of an [event
handler IDL
attribute](webappapis.html#event-handler-idl-attributes){#shared-workers-introduction:event-handler-idl-attributes},
and second, a message is sent *to* the worker, causing the worker to
send another message in return. Received messages are again displayed in
a log.

Here is the HTML page:

``` html
<!DOCTYPE HTML>
<html lang="en">
<meta charset="utf-8">
<title>Shared workers: demo 2</title>
<pre id="log">Log:</pre>
<script>
  var worker = new SharedWorker('test.js');
  var log = document.getElementById('log');
  worker.port.addEventListener('message', function(e) {
    log.textContent += '\n' + e.data;
  }, false);
  worker.port.start(); // note: need this when using addEventListener
  worker.port.postMessage('ping');
</script>
```

Here is the JavaScript worker:

``` js
onconnect = function(e) {
  var port = e.ports[0];
  port.postMessage('Hello World!');
  port.onmessage = function(e) {
    port.postMessage('pong'); // not e.ports[0].postMessage!
    // e.target.postMessage('pong'); would work also
  }
}
```

[View this example online](/demos/workers/shared/002/test.html).

------------------------------------------------------------------------

Finally, the example is extended to show how two pages can connect to
the same worker; in this case, the second page is merely in an
[`iframe`{#shared-workers-introduction:the-iframe-element}](iframe-embed-object.html#the-iframe-element)
on the first page, but the same principle would apply to an entirely
separate page in a separate [top-level
traversable](document-sequences.html#top-level-traversable){#shared-workers-introduction:top-level-traversable}.

Here is the outer HTML page:

``` html
<!DOCTYPE HTML>
<html lang="en">
<meta charset="utf-8">
<title>Shared workers: demo 3</title>
<pre id="log">Log:</pre>
<script>
  var worker = new SharedWorker('test.js');
  var log = document.getElementById('log');
  worker.port.addEventListener('message', function(e) {
    log.textContent += '\n' + e.data;
  }, false);
  worker.port.start();
  worker.port.postMessage('ping');
</script>
<iframe src="inner.html"></iframe>
```

Here is the inner HTML page:

``` html
<!DOCTYPE HTML>
<html lang="en">
<meta charset="utf-8">
<title>Shared workers: demo 3 inner frame</title>
<pre id=log>Inner log:</pre>
<script>
  var worker = new SharedWorker('test.js');
  var log = document.getElementById('log');
  worker.port.onmessage = function(e) {
   log.textContent += '\n' + e.data;
  }
</script>
```

Here is the JavaScript worker:

``` js
var count = 0;
onconnect = function(e) {
  count += 1;
  var port = e.ports[0];
  port.postMessage('Hello World! You are connection #' + count);
  port.onmessage = function(e) {
    port.postMessage('pong');
  }
}
```

[View this example online](/demos/workers/shared/003/test.html).

##### [10.1.2.4]{.secno} Shared state using a shared worker[](#shared-state-using-a-shared-worker){.self-link}

*This section is non-normative.*

In this example, multiple windows (viewers) can be opened that are all
viewing the same map. All the windows share the same map information,
with a single worker coordinating all the viewers. Each viewer can move
around independently, but if they set any data on the map, all the
viewers are updated.

The main page isn\'t interesting, it merely provides a way to open the
viewers:

``` html
<!DOCTYPE HTML>
<html lang="en">
 <head>
  <meta charset="utf-8">
  <title>Workers example: Multiviewer</title>
  <script>
   function openViewer() {
     window.open('viewer.html');
   }
  </script>
 </head>
 <body>
  <p><button type=button onclick="openViewer()">Open a new
  viewer</button></p>
  <p>Each viewer opens in a new window. You can have as many viewers
  as you like, they all view the same data.</p>
 </body>
</html>
```

The viewer is more involved:

``` html
<!DOCTYPE HTML>
<html lang="en">
 <head>
  <meta charset="utf-8">
  <title>Workers example: Multiviewer viewer</title>
  <script>
   var worker = new SharedWorker('worker.js', 'core');

   // CONFIGURATION
   function configure(event) {
     if (event.data.substr(0, 4) != 'cfg ') return;
     var name = event.data.substr(4).split(' ', 1)[0];
     // update display to mention our name is name
     document.getElementsByTagName('h1')[0].textContent += ' ' + name;
     // no longer need this listener
     worker.port.removeEventListener('message', configure, false);
   }
   worker.port.addEventListener('message', configure, false);

   // MAP
   function paintMap(event) {
     if (event.data.substr(0, 4) != 'map ') return;
     var data = event.data.substr(4).split(',');
     // display tiles data[0] .. data[8]
     var canvas = document.getElementById('map');
     var context = canvas.getContext('2d');
     for (var y = 0; y < 3; y += 1) {
       for (var x = 0; x < 3; x += 1) {
         var tile = data[y * 3 + x];
         if (tile == '0')
           context.fillStyle = 'green';
         else
           context.fillStyle = 'maroon';
         context.fillRect(x * 50, y * 50, 50, 50);
       }
     }
   }
   worker.port.addEventListener('message', paintMap, false);

   // PUBLIC CHAT
   function updatePublicChat(event) {
     if (event.data.substr(0, 4) != 'txt ') return;
     var name = event.data.substr(4).split(' ', 1)[0];
     var message = event.data.substr(4 + name.length + 1);
     // display "<name> message" in public chat
     var public = document.getElementById('public');
     var p = document.createElement('p');
     var n = document.createElement('button');
     n.textContent = '<' + name + '> ';
     n.onclick = function () { worker.port.postMessage('msg ' + name); };
     p.appendChild(n);
     var m = document.createElement('span');
     m.textContent = message;
     p.appendChild(m);
     public.appendChild(p);
   }
   worker.port.addEventListener('message', updatePublicChat, false);

   // PRIVATE CHAT
   function startPrivateChat(event) {
     if (event.data.substr(0, 4) != 'msg ') return;
     var name = event.data.substr(4).split(' ', 1)[0];
     var port = event.ports[0];
     // display a private chat UI
     var ul = document.getElementById('private');
     var li = document.createElement('li');
     var h3 = document.createElement('h3');
     h3.textContent = 'Private chat with ' + name;
     li.appendChild(h3);
     var div = document.createElement('div');
     var addMessage = function(name, message) {
       var p = document.createElement('p');
       var n = document.createElement('strong');
       n.textContent = '<' + name + '> ';
       p.appendChild(n);
       var t = document.createElement('span');
       t.textContent = message;
       p.appendChild(t);
       div.appendChild(p);
     };
     port.onmessage = function (event) {
       addMessage(name, event.data);
     };
     li.appendChild(div);
     var form = document.createElement('form');
     var p = document.createElement('p');
     var input = document.createElement('input');
     input.size = 50;
     p.appendChild(input);
     p.appendChild(document.createTextNode(' '));
     var button = document.createElement('button');
     button.textContent = 'Post';
     p.appendChild(button);
     form.onsubmit = function () {
       port.postMessage(input.value);
       addMessage('me', input.value);
       input.value = '';
       return false;
     };
     form.appendChild(p);
     li.appendChild(form);
     ul.appendChild(li);
   }
   worker.port.addEventListener('message', startPrivateChat, false);

   worker.port.start();
  </script>
 </head>
 <body>
  <h1>Viewer</h1>
  <h2>Map</h2>
  <p><canvas id="map" height=150 width=150></canvas></p>
  <p>
   <button type=button onclick="worker.port.postMessage('mov left')">Left</button>
   <button type=button onclick="worker.port.postMessage('mov up')">Up</button>
   <button type=button onclick="worker.port.postMessage('mov down')">Down</button>
   <button type=button onclick="worker.port.postMessage('mov right')">Right</button>
   <button type=button onclick="worker.port.postMessage('set 0')">Set 0</button>
   <button type=button onclick="worker.port.postMessage('set 1')">Set 1</button>
  </p>
  <h2>Public Chat</h2>
  <div id="public"></div>
  <form onsubmit="worker.port.postMessage('txt ' + message.value); message.value = ''; return false;">
   <p>
    <input type="text" name="message" size="50">
    <button>Post</button>
   </p>
  </form>
  <h2>Private Chat</h2>
  <ul id="private"></ul>
 </body>
</html>
```

There are several key things worth noting about the way the viewer is
written.

**Multiple listeners**. Instead of a single message processing function,
the code here attaches multiple event listeners, each one performing a
quick check to see if it is relevant for the message. In this example it
doesn\'t make much difference, but if multiple authors wanted to
collaborate using a single port to communicate with a worker, it would
allow for independent code instead of changes having to all be made to a
single event handling function.

Registering event listeners in this way also allows you to unregister
specific listeners when you are done with them, as is done with the
`configure()` method in this example.

Finally, the worker:

``` js
var nextName = 0;
function getNextName() {
  // this could use more friendly names
  // but for now just return a number
  return nextName++;
}

var map = [
 [0, 0, 0, 0, 0, 0, 0],
 [1, 1, 0, 1, 0, 1, 1],
 [0, 1, 0, 1, 0, 0, 0],
 [0, 1, 0, 1, 0, 1, 1],
 [0, 0, 0, 1, 0, 0, 0],
 [1, 0, 0, 1, 1, 1, 1],
 [1, 1, 0, 1, 1, 0, 1],
];

function wrapX(x) {
  if (x < 0) return wrapX(x + map[0].length);
  if (x >= map[0].length) return wrapX(x - map[0].length);
  return x;
}

function wrapY(y) {
  if (y < 0) return wrapY(y + map.length);
  if (y >= map[0].length) return wrapY(y - map.length);
  return y;
}

function wrap(val, min, max) {
  if (val < min)
    return val + (max-min)+1;
  if (val > max)
    return val - (max-min)-1;
  return val;
}

function sendMapData(viewer) {
  var data = '';
  for (var y = viewer.y-1; y <= viewer.y+1; y += 1) {
    for (var x = viewer.x-1; x <= viewer.x+1; x += 1) {
      if (data != '')
        data += ',';
      data += map[wrap(y, 0, map[0].length-1)][wrap(x, 0, map.length-1)];
    }
  }
  viewer.port.postMessage('map ' + data);
}

var viewers = {};
onconnect = function (event) {
  var name = getNextName();
  event.ports[0]._data = { port: event.ports[0], name: name, x: 0, y: 0, };
  viewers[name] = event.ports[0]._data;
  event.ports[0].postMessage('cfg ' + name);
  event.ports[0].onmessage = getMessage;
  sendMapData(event.ports[0]._data);
};

function getMessage(event) {
  switch (event.data.substr(0, 4)) {
    case 'mov ':
      var direction = event.data.substr(4);
      var dx = 0;
      var dy = 0;
      switch (direction) {
        case 'up': dy = -1; break;
        case 'down': dy = 1; break;
        case 'left': dx = -1; break;
        case 'right': dx = 1; break;
      }
      event.target._data.x = wrapX(event.target._data.x + dx);
      event.target._data.y = wrapY(event.target._data.y + dy);
      sendMapData(event.target._data);
      break;
    case 'set ':
      var value = event.data.substr(4);
      map[event.target._data.y][event.target._data.x] = value;
      for (var viewer in viewers)
        sendMapData(viewers[viewer]);
      break;
    case 'txt ':
      var name = event.target._data.name;
      var message = event.data.substr(4);
      for (var viewer in viewers)
        viewers[viewer].port.postMessage('txt ' + name + ' ' + message);
      break;
    case 'msg ':
      var party1 = event.target._data;
      var party2 = viewers[event.data.substr(4).split(' ', 1)[0]];
      if (party2) {
        var channel = new MessageChannel();
        party1.port.postMessage('msg ' + party2.name, [channel.port1]);
        party2.port.postMessage('msg ' + party1.name, [channel.port2]);
      }
      break;
  }
}
```

**Connecting to multiple pages**. The script uses the
[`onconnect`{#shared-state-using-a-shared-worker:handler-sharedworkerglobalscope-onconnect}](#handler-sharedworkerglobalscope-onconnect)
event listener to listen for multiple connections.

**Direct channels**. When the worker receives a \"msg\" message from one
viewer naming another viewer, it sets up a direct connection between the
two, so that the two viewers can communicate directly without the worker
having to proxy all the messages.

[View this example online](/demos/workers/multiviewer/page.html).

##### [10.1.2.5]{.secno} Delegation[](#delegation){.self-link}

*This section is non-normative.*

With multicore CPUs becoming prevalent, one way to obtain better
performance is to split computationally expensive tasks amongst multiple
workers. In this example, a computationally expensive task that is to be
performed for every number from 1 to 10,000,000 is farmed out to ten
subworkers.

The main page is as follows, it just reports the result:

``` html
<!DOCTYPE HTML>
<html lang="en">
 <head>
  <meta charset="utf-8">
  <title>Worker example: Multicore computation</title>
 </head>
 <body>
  <p>Result: <output id="result"></output></p>
  <script>
   var worker = new Worker('worker.js');
   worker.onmessage = function (event) {
     document.getElementById('result').textContent = event.data;
   };
  </script>
 </body>
</html>
```

The worker itself is as follows:

``` js
// settings
var num_workers = 10;
var items_per_worker = 1000000;

// start the workers
var result = 0;
var pending_workers = num_workers;
for (var i = 0; i < num_workers; i += 1) {
  var worker = new Worker('core.js');
  worker.postMessage(i * items_per_worker);
  worker.postMessage((i+1) * items_per_worker);
  worker.onmessage = storeResult;
}

// handle the results
function storeResult(event) {
  result += 1*event.data;
  pending_workers -= 1;
  if (pending_workers <= 0)
    postMessage(result); // finished!
}
```

It consists of a loop to start the subworkers, and then a handler that
waits for all the subworkers to respond.

The subworkers are implemented as follows:

``` js
var start;
onmessage = getStart;
function getStart(event) {
  start = 1*event.data;
  onmessage = getEnd;
}

var end;
function getEnd(event) {
  end = 1*event.data;
  onmessage = null;
  work();
}

function work() {
  var result = 0;
  for (var i = start; i < end; i += 1) {
    // perform some complex calculation here
    result += 1;
  }
  postMessage(result);
  close();
}
```

They receive two numbers in two events, perform the computation for the
range of numbers thus specified, and then report the result back to the
parent.

[View this example online](/demos/workers/multicore/page.html).

##### [10.1.2.6]{.secno} Providing libraries[](#providing-libraries){.self-link}

*This section is non-normative.*

Suppose that a cryptography library is made available that provides
three tasks:

Generate a public/private key pair
:   Takes a port, on which it will send two messages, first the public
    key and then the private key.

Given a plaintext and a public key, return the corresponding ciphertext
:   Takes a port, to which any number of messages can be sent, the first
    giving the public key, and the remainder giving the plaintext, each
    of which is encrypted and then sent on that same channel as the
    ciphertext. The user can close the port when it is done encrypting
    content.

Given a ciphertext and a private key, return the corresponding plaintext
:   Takes a port, to which any number of messages can be sent, the first
    giving the private key, and the remainder giving the ciphertext,
    each of which is decrypted and then sent on that same channel as the
    plaintext. The user can close the port when it is done decrypting
    content.

The library itself is as follows:

``` js
function handleMessage(e) {
  if (e.data == "genkeys")
    genkeys(e.ports[0]);
  else if (e.data == "encrypt")
    encrypt(e.ports[0]);
  else if (e.data == "decrypt")
    decrypt(e.ports[0]);
}

function genkeys(p) {
  var keys = _generateKeyPair();
  p.postMessage(keys[0]);
  p.postMessage(keys[1]);
}

function encrypt(p) {
  var key, state = 0;
  p.onmessage = function (e) {
    if (state == 0) {
      key = e.data;
      state = 1;
    } else {
      p.postMessage(_encrypt(key, e.data));
    }
  };
}

function decrypt(p) {
  var key, state = 0;
  p.onmessage = function (e) {
    if (state == 0) {
      key = e.data;
      state = 1;
    } else {
      p.postMessage(_decrypt(key, e.data));
    }
  };
}

// support being used as a shared worker as well as a dedicated worker
if ('onmessage' in this) // dedicated worker
  onmessage = handleMessage;
else // shared worker
  onconnect = function (e) { e.port.onmessage = handleMessage; }


// the "crypto" functions:

function _generateKeyPair() {
  return [Math.random(), Math.random()];
}

function _encrypt(k, s) {
  return 'encrypted-' + k + ' ' + s;
}

function _decrypt(k, s) {
  return s.substr(s.indexOf(' ')+1);
}
```

Note that the crypto functions here are just stubs and don\'t do real
cryptography.

This library could be used as follows:

``` html
<!DOCTYPE HTML>
<html lang="en">
 <head>
  <meta charset="utf-8">
  <title>Worker example: Crypto library</title>
  <script>
   const cryptoLib = new Worker('libcrypto-v1.js'); // or could use 'libcrypto-v2.js'
   function startConversation(source, message) {
     const messageChannel = new MessageChannel();
     source.postMessage(message, [messageChannel.port2]);
     return messageChannel.port1;
   }
   function getKeys() {
     let state = 0;
     startConversation(cryptoLib, "genkeys").onmessage = function (e) {
       if (state === 0)
         document.getElementById('public').value = e.data;
       else if (state === 1)
         document.getElementById('private').value = e.data;
       state += 1;
     };
   }
   function enc() {
     const port = startConversation(cryptoLib, "encrypt");
     port.postMessage(document.getElementById('public').value);
     port.postMessage(document.getElementById('input').value);
     port.onmessage = function (e) {
       document.getElementById('input').value = e.data;
       port.close();
     };
   }
   function dec() {
     const port = startConversation(cryptoLib, "decrypt");
     port.postMessage(document.getElementById('private').value);
     port.postMessage(document.getElementById('input').value);
     port.onmessage = function (e) {
       document.getElementById('input').value = e.data;
       port.close();
     };
   }
  </script>
  <style>
   textarea { display: block; }
  </style>
 </head>
 <body onload="getKeys()">
  <fieldset>
   <legend>Keys</legend>
   <p><label>Public Key: <textarea id="public"></textarea></label></p>
   <p><label>Private Key: <textarea id="private"></textarea></label></p>
  </fieldset>
  <p><label>Input: <textarea id="input"></textarea></label></p>
  <p><button onclick="enc()">Encrypt</button> <button onclick="dec()">Decrypt</button></p>
 </body>
</html>
```

A later version of the API, though, might want to offload all the crypto
work onto subworkers. This could be done as follows:

``` js
function handleMessage(e) {
  if (e.data == "genkeys")
    genkeys(e.ports[0]);
  else if (e.data == "encrypt")
    encrypt(e.ports[0]);
  else if (e.data == "decrypt")
    decrypt(e.ports[0]);
}

function genkeys(p) {
  var generator = new Worker('libcrypto-v2-generator.js');
  generator.postMessage('', [p]);
}

function encrypt(p) {
  p.onmessage = function (e) {
    var key = e.data;
    var encryptor = new Worker('libcrypto-v2-encryptor.js');
    encryptor.postMessage(key, [p]);
  };
}

function encrypt(p) {
  p.onmessage = function (e) {
    var key = e.data;
    var decryptor = new Worker('libcrypto-v2-decryptor.js');
    decryptor.postMessage(key, [p]);
  };
}

// support being used as a shared worker as well as a dedicated worker
if ('onmessage' in this) // dedicated worker
  onmessage = handleMessage;
else // shared worker
  onconnect = function (e) { e.ports[0].onmessage = handleMessage };
```

The little subworkers would then be as follows.

For generating key pairs:

``` js
onmessage = function (e) {
  var k = _generateKeyPair();
  e.ports[0].postMessage(k[0]);
  e.ports[0].postMessage(k[1]);
  close();
}

function _generateKeyPair() {
  return [Math.random(), Math.random()];
}
```

For encrypting:

``` js
onmessage = function (e) {
  var key = e.data;
  e.ports[0].onmessage = function (e) {
    var s = e.data;
    postMessage(_encrypt(key, s));
  }
}

function _encrypt(k, s) {
  return 'encrypted-' + k + ' ' + s;
}
```

For decrypting:

``` js
onmessage = function (e) {
  var key = e.data;
  e.ports[0].onmessage = function (e) {
    var s = e.data;
    postMessage(_decrypt(key, s));
  }
}

function _decrypt(k, s) {
  return s.substr(s.indexOf(' ')+1);
}
```

Notice how the users of the API don\'t have to even know that this is
happening --- the API hasn\'t changed; the library can delegate to
subworkers without changing its API, even though it is accepting data
using message channels.

[View this example online](/demos/workers/crypto/page.html).

#### [10.1.3]{.secno} Tutorials[](#tutorials){.self-link}

##### [10.1.3.1]{.secno} Creating a dedicated worker[](#creating-a-dedicated-worker){.self-link}

*This section is non-normative.*

Creating a worker requires a URL to a JavaScript file. The
[`Worker()`{#creating-a-dedicated-worker:dom-worker}](#dom-worker)
constructor is invoked with the URL to that file as its only argument; a
worker is then created and returned:

``` js
var worker = new Worker('helper.js');
```

If you want your worker script to be interpreted as a [module
script](webappapis.html#module-script){#creating-a-dedicated-worker:module-script}
instead of the default [classic
script](webappapis.html#classic-script){#creating-a-dedicated-worker:classic-script},
you need to use a slightly different signature:

``` js
var worker = new Worker('helper.mjs', { type: "module" });
```

##### [10.1.3.2]{.secno} Communicating with a dedicated worker[](#communicating-with-a-dedicated-worker){.self-link}

*This section is non-normative.*

Dedicated workers use
[`MessagePort`{#communicating-with-a-dedicated-worker:messageport}](web-messaging.html#messageport)
objects behind the scenes, and thus support all the same features, such
as sending structured data, transferring binary data, and transferring
other ports.

To receive messages from a dedicated worker, use the
[`onmessage`{#communicating-with-a-dedicated-worker:handler-messageeventtarget-onmessage}](web-messaging.html#handler-messageeventtarget-onmessage)
[event handler IDL
attribute](webappapis.html#event-handler-idl-attributes){#communicating-with-a-dedicated-worker:event-handler-idl-attributes}
on the
[`Worker`{#communicating-with-a-dedicated-worker:worker}](#worker)
object:

``` js
worker.onmessage = function (event) { ... };
```

You can also use the
[`addEventListener()`{#communicating-with-a-dedicated-worker:dom-eventtarget-addeventlistener}](https://dom.spec.whatwg.org/#dom-eventtarget-addeventlistener){x-internal="dom-eventtarget-addeventlistener"}
method.

The implicit
[`MessagePort`{#communicating-with-a-dedicated-worker:messageport-2}](web-messaging.html#messageport)
used by dedicated workers has its [port message
queue](web-messaging.html#port-message-queue){#communicating-with-a-dedicated-worker:port-message-queue}
implicitly enabled when it is created, so there is no equivalent to the
[`MessagePort`{#communicating-with-a-dedicated-worker:messageport-3}](web-messaging.html#messageport)
interface\'s
[`start()`{#communicating-with-a-dedicated-worker:dom-messageport-start}](web-messaging.html#dom-messageport-start)
method on the
[`Worker`{#communicating-with-a-dedicated-worker:worker-2}](#worker)
interface.

To *send* data to a worker, use the
[`postMessage()`{#communicating-with-a-dedicated-worker:dom-worker-postmessage}](#dom-worker-postmessage)
method. Structured data can be sent over this communication channel. To
send
[`ArrayBuffer`{#communicating-with-a-dedicated-worker:idl-arraybuffer}](https://webidl.spec.whatwg.org/#idl-ArrayBuffer){x-internal="idl-arraybuffer"}
objects efficiently (by transferring them rather than cloning them),
list them in an array in the second argument.

``` js
worker.postMessage({
  operation: 'find-edges',
  input: buffer, // an ArrayBuffer object
  threshold: 0.6,
}, [buffer]);
```

To receive a message inside the worker, the
[`onmessage`{#communicating-with-a-dedicated-worker:handler-messageeventtarget-onmessage-2}](web-messaging.html#handler-messageeventtarget-onmessage)
[event handler IDL
attribute](webappapis.html#event-handler-idl-attributes){#communicating-with-a-dedicated-worker:event-handler-idl-attributes-2}
is used.

``` js
onmessage = function (event) { ... };
```

You can again also use the
[`addEventListener()`{#communicating-with-a-dedicated-worker:dom-eventtarget-addeventlistener-2}](https://dom.spec.whatwg.org/#dom-eventtarget-addeventlistener){x-internal="dom-eventtarget-addeventlistener"}
method.

In either case, the data is provided in the event object\'s
[`data`{#communicating-with-a-dedicated-worker:dom-messageevent-data}](comms.html#dom-messageevent-data)
attribute.

To send messages back, you again use
[`postMessage()`{#communicating-with-a-dedicated-worker:dom-dedicatedworkerglobalscope-postmessage}](#dom-dedicatedworkerglobalscope-postmessage).
It supports the structured data in the same manner.

``` js
postMessage(event.data.input, [event.data.input]); // transfer the buffer back
```

##### [10.1.3.3]{.secno} Shared workers[](#shared-workers){.self-link} {#shared-workers dfn-type="dfn" lt="shared worker"}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[SharedWorker](https://developer.mozilla.org/en-US/docs/Web/API/SharedWorker "The SharedWorker interface represents a specific kind of worker that can be accessed from several browsing contexts, such as several windows, iframes or even workers. They implement an interface different than dedicated workers and have a different global scope, SharedWorkerGlobalScope.")

Support in all current engines.

::: support
[Firefox29+]{.firefox .yes}[Safari16+]{.safari .yes}[Chrome5+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android33+]{.firefox_android .yes}[Safari iOS16+]{.safari_ios
.yes}[Chrome AndroidNo]{.chrome_android .no}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet4.0--5.0]{.samsunginternet_android .no}[Opera
Android11--14]{.opera_android .no}
:::
::::
:::::

*This section is non-normative.*

Shared workers are identified by the URL of the script used to create
it, optionally with an explicit name. The name allows multiple instances
of a particular shared worker to be started.

Shared workers are scoped by
[origin](browsers.html#concept-origin){#shared-workers:concept-origin}.
Two different sites using the same names will not collide. However, if a
page tries to use the same shared worker name as another page on the
same site, but with a different script URL, it will fail.

Creating shared workers is done using the
[`SharedWorker()`{#shared-workers:dom-sharedworker}](#dom-sharedworker)
constructor. This constructor takes the URL to the script to use for its
first argument, and the name of the worker, if any, as the second
argument.

``` js
var worker = new SharedWorker('service.js');
```

Communicating with shared workers is done with explicit
[`MessagePort`{#shared-workers:messageport}](web-messaging.html#messageport)
objects. The object returned by the
[`SharedWorker()`{#shared-workers:dom-sharedworker-2}](#dom-sharedworker)
constructor holds a reference to the port on its
[`port`{#shared-workers:dom-sharedworker-port}](#dom-sharedworker-port)
attribute.

``` js
worker.port.onmessage = function (event) { ... };
worker.port.postMessage('some message');
worker.port.postMessage({ foo: 'structured', bar: ['data', 'also', 'possible']});
```

Inside the shared worker, new clients of the worker are announced using
the
[`connect`{#shared-workers:event-workerglobalscope-connect}](indices.html#event-workerglobalscope-connect)
event. The port for the new client is given by the event object\'s
[`source`{#shared-workers:dom-messageevent-source}](comms.html#dom-messageevent-source)
attribute.

``` js
onconnect = function (event) {
  var newPort = event.source;
  // set up a listener
  newPort.onmessage = function (event) { ... };
  // send a message back to the port
  newPort.postMessage('ready!'); // can also send structured data, of course
};
```

### [10.2]{.secno} Infrastructure[](#infrastructure-2){.self-link} {#infrastructure-2}

This standard defines two kinds of workers: dedicated workers, and
shared workers. Dedicated workers, once created, are linked to their
creator, but message ports can be used to communicate from a dedicated
worker to multiple other browsing contexts or workers. Shared workers,
on the other hand, are named, and once created any script running in the
same
[origin](browsers.html#concept-origin){#infrastructure-2:concept-origin}
can obtain a reference to that worker and communicate with it. Service
Workers defines a third kind. [\[SW\]](references.html#refsSW)

#### [10.2.1]{.secno} The global scope[](#the-global-scope){.self-link}

The global scope is the \"inside\" of a worker.

##### [10.2.1.1]{.secno} The [`WorkerGlobalScope`{#workerglobalscope-dev}](#workerglobalscope) common interface[](#the-workerglobalscope-common-interface){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[WorkerGlobalScope](https://developer.mozilla.org/en-US/docs/Web/API/WorkerGlobalScope "The WorkerGlobalScope interface of the Web Workers API is an interface representing the scope of any worker. Workers have no browsing context; this scope contains the information usually conveyed by Window objects — in this case event handlers, the console or the associated WorkerNavigator object. Each WorkerGlobalScope has its own event loop.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

``` idl
[Exposed=Worker]
interface WorkerGlobalScope : EventTarget {
  readonly attribute WorkerGlobalScope self;
  readonly attribute WorkerLocation location;
  readonly attribute WorkerNavigator navigator;
  undefined importScripts((TrustedScriptURL or USVString)... urls);

  attribute OnErrorEventHandler onerror;
  attribute EventHandler onlanguagechange;
  attribute EventHandler onoffline;
  attribute EventHandler ononline;
  attribute EventHandler onrejectionhandled;
  attribute EventHandler onunhandledrejection;
};
```

[`WorkerGlobalScope`{#the-workerglobalscope-common-interface:workerglobalscope-2}](#workerglobalscope)
serves as the base class for specific types of worker global scope
objects, including
[`DedicatedWorkerGlobalScope`{#the-workerglobalscope-common-interface:dedicatedworkerglobalscope}](#dedicatedworkerglobalscope),
[`SharedWorkerGlobalScope`{#the-workerglobalscope-common-interface:sharedworkerglobalscope}](#sharedworkerglobalscope),
and
[`ServiceWorkerGlobalScope`{#the-workerglobalscope-common-interface:serviceworkerglobalscope}](https://w3c.github.io/ServiceWorker/#serviceworkerglobalscope){x-internal="serviceworkerglobalscope"}.

A
[`WorkerGlobalScope`{#the-workerglobalscope-common-interface:workerglobalscope-3}](#workerglobalscope)
object has an associated [owner
set]{#concept-WorkerGlobalScope-owner-set .dfn
dfn-for="WorkerGlobalScope" export=""} (a
[set](https://infra.spec.whatwg.org/#ordered-set){#the-workerglobalscope-common-interface:set
x-internal="set"} of
[`Document`{#the-workerglobalscope-common-interface:document}](dom.html#document)
and
[`WorkerGlobalScope`{#the-workerglobalscope-common-interface:workerglobalscope-4}](#workerglobalscope)
objects). It is initially empty and populated when the worker is created
or obtained.

It is a
[set](https://infra.spec.whatwg.org/#ordered-set){#the-workerglobalscope-common-interface:set-2
x-internal="set"}, instead of a single owner, to accommodate
[`SharedWorkerGlobalScope`{#the-workerglobalscope-common-interface:sharedworkerglobalscope-2}](#sharedworkerglobalscope)
objects.

A
[`WorkerGlobalScope`{#the-workerglobalscope-common-interface:workerglobalscope-5}](#workerglobalscope)
object has an associated [type]{#concept-workerglobalscope-type .dfn
dfn-for="WorkerGlobalScope" export=""} (\"`classic`\" or \"`module`\").
It is set during creation.

A
[`WorkerGlobalScope`{#the-workerglobalscope-common-interface:workerglobalscope-6}](#workerglobalscope)
object has an associated [url]{#concept-workerglobalscope-url .dfn
dfn-for="WorkerGlobalScope" export=""} (null or a
[URL](https://url.spec.whatwg.org/#concept-url){#the-workerglobalscope-common-interface:url
x-internal="url"}). It is initially null.

A
[`WorkerGlobalScope`{#the-workerglobalscope-common-interface:workerglobalscope-7}](#workerglobalscope)
object has an associated [name]{#concept-workerglobalscope-name .dfn
dfn-for="WorkerGlobalScope" export=""} (a string). It is set during
creation.

The
[name](#concept-workerglobalscope-name){#the-workerglobalscope-common-interface:concept-workerglobalscope-name}
can have different semantics for each subclass of
[`WorkerGlobalScope`{#the-workerglobalscope-common-interface:workerglobalscope-8}](#workerglobalscope).
For
[`DedicatedWorkerGlobalScope`{#the-workerglobalscope-common-interface:dedicatedworkerglobalscope-2}](#dedicatedworkerglobalscope)
instances, it is simply a developer-supplied name, useful mostly for
debugging purposes. For
[`SharedWorkerGlobalScope`{#the-workerglobalscope-common-interface:sharedworkerglobalscope-3}](#sharedworkerglobalscope)
instances, it allows obtaining a reference to a common shared worker via
the
[`SharedWorker()`{#the-workerglobalscope-common-interface:dom-sharedworker}](#dom-sharedworker)
constructor. For
[`ServiceWorkerGlobalScope`{#the-workerglobalscope-common-interface:serviceworkerglobalscope-2}](https://w3c.github.io/ServiceWorker/#serviceworkerglobalscope){x-internal="serviceworkerglobalscope"}
objects, it doesn\'t make sense (and as such isn\'t exposed through the
JavaScript API at all).

A
[`WorkerGlobalScope`{#the-workerglobalscope-common-interface:workerglobalscope-9}](#workerglobalscope)
object has an associated [policy
container]{#concept-workerglobalscope-policy-container .dfn
dfn-for="WorkerGlobalScope" export=""} (a [policy
container](browsers.html#policy-container){#the-workerglobalscope-common-interface:policy-container}).
It is initially a new [policy
container](browsers.html#policy-container){#the-workerglobalscope-common-interface:policy-container-2}.

A
[`WorkerGlobalScope`{#the-workerglobalscope-common-interface:workerglobalscope-10}](#workerglobalscope)
object has an associated [embedder
policy]{#concept-workerglobalscope-embedder-policy .dfn
dfn-for="WorkerGlobalScope" export=""} (an [embedder
policy](browsers.html#embedder-policy){#the-workerglobalscope-common-interface:embedder-policy}).

A
[`WorkerGlobalScope`{#the-workerglobalscope-common-interface:workerglobalscope-11}](#workerglobalscope)
object has an associated [module
map]{#concept-workerglobalscope-module-map .dfn
dfn-for="WorkerGlobalScope" export=""}. It is a [module
map](webappapis.html#module-map){#the-workerglobalscope-common-interface:module-map},
initially empty.

A
[`WorkerGlobalScope`{#the-workerglobalscope-common-interface:workerglobalscope-12}](#workerglobalscope)
object has an associated [cross-origin isolated
capability]{#concept-workerglobalscope-cross-origin-isolated-capability
.dfn dfn-for="WorkerGlobalScope" export=""} boolean. It is initially
false.

`workerGlobal`{.variable}`.`[`self`](#dom-workerglobalscope-self){#dom-workerglobalscope-self-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[WorkerGlobalScope/self](https://developer.mozilla.org/en-US/docs/Web/API/WorkerGlobalScope/self "The self read-only property of the WorkerGlobalScope interface returns a reference to the WorkerGlobalScope itself. Most of the time it is a specific scope like DedicatedWorkerGlobalScope, SharedWorkerGlobalScope, or ServiceWorkerGlobalScope.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android34+]{.firefox_android .yes}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

Returns `workerGlobal`{.variable}.

`workerGlobal`{.variable}`.`[`location`](#dom-workerglobalscope-location){#dom-workerglobalscope-location-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[WorkerGlobalScope/location](https://developer.mozilla.org/en-US/docs/Web/API/WorkerGlobalScope/location "The location read-only property of the WorkerGlobalScope interface returns the WorkerLocation associated with the worker. It is a specific location object, mostly a subset of the Location for browsing scopes, but adapted to workers.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

Returns `workerGlobal`{.variable}\'s
[`WorkerLocation`{#the-workerglobalscope-common-interface:workerlocation-2}](#workerlocation)
object.

`workerGlobal`{.variable}`.`[`navigator`](#dom-worker-navigator){#dom-worker-navigator-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[WorkerGlobalScope/navigator](https://developer.mozilla.org/en-US/docs/Web/API/WorkerGlobalScope/navigator "The navigator read-only property of the WorkerGlobalScope interface returns the WorkerNavigator associated with the worker. It is a specific navigator object, mostly a subset of the Navigator for browsing scopes, but adapted to workers.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)17+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

Returns `workerGlobal`{.variable}\'s
[`WorkerNavigator`{#the-workerglobalscope-common-interface:workernavigator-2}](#workernavigator)
object.

`workerGlobal`{.variable}`.`[`importScripts`](#dom-workerglobalscope-importscripts){#dom-workerglobalscope-importscripts-dev}`(...``urls`{.variable}`)`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[WorkerGlobalScope/importScripts](https://developer.mozilla.org/en-US/docs/Web/API/WorkerGlobalScope/importScripts "The importScripts() method of the WorkerGlobalScope interface synchronously imports one or more scripts into the worker's scope.")

Support in all current engines.

::: support
[Firefox4+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

Fetches each
[URL](https://url.spec.whatwg.org/#concept-url){#the-workerglobalscope-common-interface:url-2
x-internal="url"} in `urls`{.variable}, executes them one-by-one in the
order they are passed, and then returns (or throws if something went
amiss).

The [`self`]{#dom-workerglobalscope-self .dfn
dfn-for="WorkerGlobalScope" dfn-type="attribute"} attribute must return
the
[`WorkerGlobalScope`{#the-workerglobalscope-common-interface:workerglobalscope-13}](#workerglobalscope)
object itself.

The [`location`]{#dom-workerglobalscope-location .dfn
dfn-for="WorkerGlobalScope" dfn-type="attribute"} attribute must return
the
[`WorkerLocation`{#the-workerglobalscope-common-interface:workerlocation-3}](#workerlocation)
object whose associated [`WorkerGlobalScope`
object](#concept-workerlocation-workerglobalscope){#the-workerglobalscope-common-interface:concept-workerlocation-workerglobalscope}
is the
[`WorkerGlobalScope`{#the-workerglobalscope-common-interface:workerglobalscope-14}](#workerglobalscope)
object.

While the
[`WorkerLocation`{#the-workerglobalscope-common-interface:workerlocation-4}](#workerlocation)
object is created after the
[`WorkerGlobalScope`{#the-workerglobalscope-common-interface:workerglobalscope-15}](#workerglobalscope)
object, this is not problematic as it cannot be observed from script.

------------------------------------------------------------------------

The following are the [event
handlers](webappapis.html#event-handlers){#the-workerglobalscope-common-interface:event-handlers}
(and their corresponding [event handler event
types](webappapis.html#event-handler-event-type){#the-workerglobalscope-common-interface:event-handler-event-type})
that must be supported, as [event handler IDL
attributes](webappapis.html#event-handler-idl-attributes){#the-workerglobalscope-common-interface:event-handler-idl-attributes},
by objects implementing the
[`WorkerGlobalScope`{#the-workerglobalscope-common-interface:workerglobalscope-16}](#workerglobalscope)
interface:

[Event
handler](webappapis.html#event-handlers){#the-workerglobalscope-common-interface:event-handlers-2}

[Event handler event
type](webappapis.html#event-handler-event-type){#the-workerglobalscope-common-interface:event-handler-event-type-2}

[`onerror`]{#handler-workerglobalscope-onerror .dfn
dfn-for="WorkerGlobalScope" dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[WorkerGlobalScope/error_event](https://developer.mozilla.org/en-US/docs/Web/API/WorkerGlobalScope/error_event "The error event of the WorkerGlobalScope interface fires when an error occurs in the worker.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

[`error`{#the-workerglobalscope-common-interface:event-error}](indices.html#event-error)

[`onlanguagechange`]{#handler-workerglobalscope-onlanguagechange .dfn
dfn-for="WorkerGlobalScope" dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[WorkerGlobalScope/languagechange_event](https://developer.mozilla.org/en-US/docs/Web/API/WorkerGlobalScope/languagechange_event "The languagechange event is fired at the global scope object when the user's preferred language changes.")

Support in all current engines.

::: support
[Firefox74+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera11.5+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

[`languagechange`{#the-workerglobalscope-common-interface:event-languagechange}](indices.html#event-languagechange)

[`onoffline`]{#handler-workerglobalscope-onoffline .dfn
dfn-for="WorkerGlobalScope" dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
MDN

:::: feature
[WorkerGlobalScope/offline_event](https://developer.mozilla.org/en-US/docs/Web/API/WorkerGlobalScope/offline_event "The offline event of the WorkerGlobalScope fires when the device loses connection to the internet.")

::: support
[Firefox29+]{.firefox .yes}[Safari8+]{.safari .yes}[ChromeNo]{.chrome
.no}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[EdgeNo]{.edge_blink .no}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

[`offline`{#the-workerglobalscope-common-interface:event-offline}](indices.html#event-offline)

[`ononline`]{#handler-workerglobalscope-ononline .dfn
dfn-for="WorkerGlobalScope" dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
MDN

:::: feature
[WorkerGlobalScope/online_event](https://developer.mozilla.org/en-US/docs/Web/API/WorkerGlobalScope/online_event "The online event of the WorkerGlobalScope fires when the device reconnects to the internet.")

::: support
[Firefox29+]{.firefox .yes}[Safari8+]{.safari .yes}[ChromeNo]{.chrome
.no}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[EdgeNo]{.edge_blink .no}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

[`online`{#the-workerglobalscope-common-interface:event-online}](indices.html#event-online)

[`onrejectionhandled`]{#handler-workerglobalscope-onrejectionhandled
.dfn dfn-for="WorkerGlobalScope" dfn-type="attribute"}

[`rejectionhandled`{#the-workerglobalscope-common-interface:event-rejectionhandled}](indices.html#event-rejectionhandled)

[`onunhandledrejection`]{#handler-workerglobalscope-onunhandledrejection
.dfn dfn-for="WorkerGlobalScope" dfn-type="attribute"}

[`unhandledrejection`{#the-workerglobalscope-common-interface:event-unhandledrejection}](indices.html#event-unhandledrejection)

##### [10.2.1.2]{.secno} Dedicated workers and the [`DedicatedWorkerGlobalScope`{#dedicatedworkerglobalscope-dev}](#dedicatedworkerglobalscope) interface[](#dedicated-workers-and-the-dedicatedworkerglobalscope-interface){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[DedicatedWorkerGlobalScope](https://developer.mozilla.org/en-US/docs/Web/API/DedicatedWorkerGlobalScope "The DedicatedWorkerGlobalScope object (the Worker global scope) is accessible through the self keyword. Some additional global functions, namespaces objects, and constructors, not typically associated with the worker global scope, but available on it, are listed in the JavaScript Reference. See also: Functions available to workers.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

``` idl
[Global=(Worker,DedicatedWorker),Exposed=DedicatedWorker]
interface DedicatedWorkerGlobalScope : WorkerGlobalScope {
  [Replaceable] readonly attribute DOMString name;

  undefined postMessage(any message, sequence<object> transfer);
  undefined postMessage(any message, optional StructuredSerializeOptions options = {});

  undefined close();
};

DedicatedWorkerGlobalScope includes MessageEventTarget;
```

[`DedicatedWorkerGlobalScope`{#dedicated-workers-and-the-dedicatedworkerglobalscope-interface:dedicatedworkerglobalscope-2}](#dedicatedworkerglobalscope)
objects have an associated [inside port]{#inside-port .dfn} (a
[`MessagePort`{#dedicated-workers-and-the-dedicatedworkerglobalscope-interface:messageport}](web-messaging.html#messageport)).
This port is part of a channel that is set up when the worker is
created, but it is not exposed. This object must never be garbage
collected before the
[`DedicatedWorkerGlobalScope`{#dedicated-workers-and-the-dedicatedworkerglobalscope-interface:dedicatedworkerglobalscope-3}](#dedicatedworkerglobalscope)
object.

`dedicatedWorkerGlobal`{.variable}`.`[`name`](#dom-dedicatedworkerglobalscope-name){#dom-dedicatedworkerglobalscope-name-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[DedicatedWorkerGlobalScope/name](https://developer.mozilla.org/en-US/docs/Web/API/DedicatedWorkerGlobalScope/name "The name read-only property of the DedicatedWorkerGlobalScope interface returns the name that the Worker was (optionally) given when it was created. This is the name that the Worker() constructor can pass to get a reference to the DedicatedWorkerGlobalScope.")

Support in all current engines.

::: support
[Firefox55+]{.firefox .yes}[Safari12.1+]{.safari
.yes}[Chrome70+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)18]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

Returns `dedicatedWorkerGlobal`{.variable}\'s
[name](#concept-workerglobalscope-name){#dedicated-workers-and-the-dedicatedworkerglobalscope-interface:concept-workerglobalscope-name},
i.e. the value given to the
[`Worker`{#dedicated-workers-and-the-dedicatedworkerglobalscope-interface:worker}](#worker)
constructor. Primarily useful for debugging.

`dedicatedWorkerGlobal`{.variable}`.`[`postMessage`](#dom-dedicatedworkerglobalscope-postmessage){#dom-dedicatedworkerglobalscope-postmessage-dev}`(``message`{.variable}` [, ``transfer`{.variable}` ])`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[DedicatedWorkerGlobalScope/postMessage](https://developer.mozilla.org/en-US/docs/Web/API/DedicatedWorkerGlobalScope/postMessage "The postMessage() method of the DedicatedWorkerGlobalScope interface sends a message to the main thread that spawned it.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

`dedicatedWorkerGlobal`{.variable}`.`[`postMessage`](#dom-dedicatedworkerglobalscope-postmessage-options){#dom-dedicatedworkerglobalscope-postmessage-options-dev}`(``message`{.variable}` [, { `[`transfer`](web-messaging.html#dom-structuredserializeoptions-transfer){#dedicated-workers-and-the-dedicatedworkerglobalscope-interface:dom-structuredserializeoptions-transfer}` } ])`

Clones `message`{.variable} and transmits it to the
[`Worker`{#dedicated-workers-and-the-dedicatedworkerglobalscope-interface:worker-2}](#worker)
object associated with `dedicatedWorkerGlobal`{.variable}.
`transfer`{.variable} can be passed as a list of objects that are to be
transferred rather than cloned.

`dedicatedWorkerGlobal`{.variable}`.`[`close`](#dom-dedicatedworkerglobalscope-close){#dom-dedicatedworkerglobalscope-close-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[DedicatedWorkerGlobalScope/close](https://developer.mozilla.org/en-US/docs/Web/API/DedicatedWorkerGlobalScope/close "The close() method of the DedicatedWorkerGlobalScope interface discards any tasks queued in the DedicatedWorkerGlobalScope's event loop, effectively closing this particular scope.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

Aborts `dedicatedWorkerGlobal`{.variable}.

The [`name`]{#dom-dedicatedworkerglobalscope-name .dfn
dfn-for="DedicatedWorkerGlobalScope" dfn-type="attribute"} getter steps
are to return
[this](https://webidl.spec.whatwg.org/#this){#dedicated-workers-and-the-dedicatedworkerglobalscope-interface:this
x-internal="this"}\'s
[name](#concept-workerglobalscope-name){#dedicated-workers-and-the-dedicatedworkerglobalscope-interface:concept-workerglobalscope-name-2}.
Its value represents the name given to the worker using the
[`Worker`{#dedicated-workers-and-the-dedicatedworkerglobalscope-interface:worker-3}](#worker)
constructor, used primarily for debugging purposes.

The
[`postMessage(``message`{.variable}`, ``transfer`{.variable}`)`]{#dom-dedicatedworkerglobalscope-postmessage
.dfn dfn-for="DedicatedWorkerGlobalScope" dfn-type="method"} and
[`postMessage(``message`{.variable}`, ``options`{.variable}`)`]{#dom-dedicatedworkerglobalscope-postmessage-options
.dfn dfn-for="DedicatedWorkerGlobalScope" dfn-type="method"} methods on
[`DedicatedWorkerGlobalScope`{#dedicated-workers-and-the-dedicatedworkerglobalscope-interface:dedicatedworkerglobalscope-4}](#dedicatedworkerglobalscope)
objects act as if, when invoked, it immediately invoked the respective
[`postMessage(`{#dedicated-workers-and-the-dedicatedworkerglobalscope-interface:dom-messageport-postmessage}`message`{.variable}`, `{#dedicated-workers-and-the-dedicatedworkerglobalscope-interface:dom-messageport-postmessage}`transfer`{.variable}`)`{#dedicated-workers-and-the-dedicatedworkerglobalscope-interface:dom-messageport-postmessage}](web-messaging.html#dom-messageport-postmessage)
and
[`postMessage(`{#dedicated-workers-and-the-dedicatedworkerglobalscope-interface:dom-messageport-postmessage-options}`message`{.variable}`, `{#dedicated-workers-and-the-dedicatedworkerglobalscope-interface:dom-messageport-postmessage-options}`options`{.variable}`)`{#dedicated-workers-and-the-dedicatedworkerglobalscope-interface:dom-messageport-postmessage-options}](web-messaging.html#dom-messageport-postmessage-options)
on the port, with the same arguments, and returned the same return
value.

To [close a worker]{#close-a-worker .dfn export=""}, given a
`workerGlobal`{.variable}, run these steps:

1.  Discard any
    [tasks](webappapis.html#concept-task){#dedicated-workers-and-the-dedicatedworkerglobalscope-interface:concept-task}
    that have been added to `workerGlobal`{.variable}\'s [relevant
    agent](webappapis.html#relevant-agent){#dedicated-workers-and-the-dedicatedworkerglobalscope-interface:relevant-agent}\'s
    [event
    loop](webappapis.html#concept-agent-event-loop){#dedicated-workers-and-the-dedicatedworkerglobalscope-interface:concept-agent-event-loop}\'s
    [task
    queues](webappapis.html#task-queue){#dedicated-workers-and-the-dedicatedworkerglobalscope-interface:task-queue}.

2.  Set `workerGlobal`{.variable}\'s
    [closing](#dom-workerglobalscope-closing){#dedicated-workers-and-the-dedicatedworkerglobalscope-interface:dom-workerglobalscope-closing}
    flag to true. (This prevents any further tasks from being queued.)

The [`close()`]{#dom-dedicatedworkerglobalscope-close .dfn
dfn-for="DedicatedWorkerGlobalScope" dfn-type="method"} method steps are
to [close a
worker](#close-a-worker){#dedicated-workers-and-the-dedicatedworkerglobalscope-interface:close-a-worker}
given
[this](https://webidl.spec.whatwg.org/#this){#dedicated-workers-and-the-dedicatedworkerglobalscope-interface:this-2
x-internal="this"}.

------------------------------------------------------------------------

##### [10.2.1.3]{.secno} Shared workers and the [`SharedWorkerGlobalScope`{#sharedworkerglobalscope-dev}](#sharedworkerglobalscope) interface[](#shared-workers-and-the-sharedworkerglobalscope-interface){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[SharedWorkerGlobalScope](https://developer.mozilla.org/en-US/docs/Web/API/SharedWorkerGlobalScope "The SharedWorkerGlobalScope object (the SharedWorker global scope) is accessible through the self keyword. Some additional global functions, namespaces objects, and constructors, not typically associated with the worker global scope, but available on it, are listed in the JavaScript Reference. See the complete list of functions available to workers.")

Support in all current engines.

::: support
[Firefox29+]{.firefox .yes}[Safari16+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS16+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

``` idl
[Global=(Worker,SharedWorker),Exposed=SharedWorker]
interface SharedWorkerGlobalScope : WorkerGlobalScope {
  [Replaceable] readonly attribute DOMString name;

  undefined close();

  attribute EventHandler onconnect;
};
```

A
[`SharedWorkerGlobalScope`{#shared-workers-and-the-sharedworkerglobalscope-interface:sharedworkerglobalscope}](#sharedworkerglobalscope)
object has an associated [constructor
origin]{#concept-sharedworkerglobalscope-constructor-origin .dfn
dfn-for="SharedWorkerGlobalScope"}, [constructor
url]{#concept-sharedworkerglobalscope-constructor-url .dfn
dfn-for="SharedWorkerGlobalScope"}, and
[credentials]{#concept-sharedworkerglobalscope-credentials .dfn
dfn-for="SharedWorkerGlobalScope"}. They are initialized when the
[`SharedWorkerGlobalScope`{#shared-workers-and-the-sharedworkerglobalscope-interface:sharedworkerglobalscope-2}](#sharedworkerglobalscope)
object is created, in the [run a
worker](#run-a-worker){#shared-workers-and-the-sharedworkerglobalscope-interface:run-a-worker}
algorithm.

Shared workers receive message ports through
[`connect`{#shared-workers-and-the-sharedworkerglobalscope-interface:event-workerglobalscope-connect}](indices.html#event-workerglobalscope-connect)
events on their
[`SharedWorkerGlobalScope`{#shared-workers-and-the-sharedworkerglobalscope-interface:sharedworkerglobalscope-3}](#sharedworkerglobalscope)
object for each connection.

`sharedWorkerGlobal`{.variable}`.`[`name`](#dom-sharedworkerglobalscope-name){#dom-sharedworkerglobalscope-name-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[SharedWorkerGlobalScope/name](https://developer.mozilla.org/en-US/docs/Web/API/SharedWorkerGlobalScope/name "The name read-only property of the SharedWorkerGlobalScope interface returns the name that the SharedWorker was (optionally) given when it was created. This is the name that the SharedWorker() constructor can pass to get a reference to the SharedWorkerGlobalScope.")

Support in all current engines.

::: support
[Firefox29+]{.firefox .yes}[Safari16+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS16+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

Returns `sharedWorkerGlobal`{.variable}\'s
[name](#concept-workerglobalscope-name){#shared-workers-and-the-sharedworkerglobalscope-interface:concept-workerglobalscope-name},
i.e. the value given to the
[`SharedWorker`{#shared-workers-and-the-sharedworkerglobalscope-interface:sharedworker}](#sharedworker)
constructor. Multiple
[`SharedWorker`{#shared-workers-and-the-sharedworkerglobalscope-interface:sharedworker-2}](#sharedworker)
objects can correspond to the same shared worker (and
[`SharedWorkerGlobalScope`{#shared-workers-and-the-sharedworkerglobalscope-interface:sharedworkerglobalscope-4}](#sharedworkerglobalscope)),
by reusing the same name.

`sharedWorkerGlobal`{.variable}`.`[`close`](#dom-sharedworkerglobalscope-close){#dom-sharedworkerglobalscope-close-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[SharedWorkerGlobalScope/close](https://developer.mozilla.org/en-US/docs/Web/API/SharedWorkerGlobalScope/close "The close() method of the SharedWorkerGlobalScope interface discards any tasks queued in the SharedWorkerGlobalScope's event loop, effectively closing this particular scope.")

Support in all current engines.

::: support
[Firefox29+]{.firefox .yes}[Safari16+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS16+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

Aborts `sharedWorkerGlobal`{.variable}.

The [`name`]{#dom-sharedworkerglobalscope-name .dfn
dfn-for="SharedWorkerGlobalScope" dfn-type="attribute"} getter steps are
to return
[this](https://webidl.spec.whatwg.org/#this){#shared-workers-and-the-sharedworkerglobalscope-interface:this
x-internal="this"}\'s
[name](#concept-workerglobalscope-name){#shared-workers-and-the-sharedworkerglobalscope-interface:concept-workerglobalscope-name-2}.
Its value represents the name that can be used to obtain a reference to
the worker using the
[`SharedWorker`{#shared-workers-and-the-sharedworkerglobalscope-interface:sharedworker-3}](#sharedworker)
constructor.

The [`close()`]{#dom-sharedworkerglobalscope-close .dfn
dfn-for="SharedWorkerGlobalScope" dfn-type="method"} method steps are to
[close a
worker](#close-a-worker){#shared-workers-and-the-sharedworkerglobalscope-interface:close-a-worker}
given
[this](https://webidl.spec.whatwg.org/#this){#shared-workers-and-the-sharedworkerglobalscope-interface:this-2
x-internal="this"}.

------------------------------------------------------------------------

The following are the [event
handlers](webappapis.html#event-handlers){#shared-workers-and-the-sharedworkerglobalscope-interface:event-handlers}
(and their corresponding [event handler event
types](webappapis.html#event-handler-event-type){#shared-workers-and-the-sharedworkerglobalscope-interface:event-handler-event-type})
that must be supported, as [event handler IDL
attributes](webappapis.html#event-handler-idl-attributes){#shared-workers-and-the-sharedworkerglobalscope-interface:event-handler-idl-attributes},
by objects implementing the
[`SharedWorkerGlobalScope`{#shared-workers-and-the-sharedworkerglobalscope-interface:sharedworkerglobalscope-5}](#sharedworkerglobalscope)
interface:

[Event
handler](webappapis.html#event-handlers){#shared-workers-and-the-sharedworkerglobalscope-interface:event-handlers-2}

[Event handler event
type](webappapis.html#event-handler-event-type){#shared-workers-and-the-sharedworkerglobalscope-interface:event-handler-event-type-2}

[`onconnect`]{#handler-sharedworkerglobalscope-onconnect .dfn
dfn-for="SharedWorkerGlobalScope" dfn-type="attribute"}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[SharedWorkerGlobalScope/connect_event](https://developer.mozilla.org/en-US/docs/Web/API/SharedWorkerGlobalScope/connect_event "The connect event is fired in shared workers at their SharedWorkerGlobalScope when a new client connects.")

Support in all current engines.

::: support
[Firefox29+]{.firefox .yes}[Safari16+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS16+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

[`connect`{#shared-workers-and-the-sharedworkerglobalscope-interface:event-workerglobalscope-connect-2}](indices.html#event-workerglobalscope-connect)

#### [10.2.2]{.secno} The event loop[](#worker-event-loop){.self-link} {#worker-event-loop}

A [worker event
loop](webappapis.html#worker-event-loop-2){#worker-event-loop:worker-event-loop-2}\'s
[task queues](webappapis.html#task-queue){#worker-event-loop:task-queue}
only have events, callbacks, and networking activity as
[tasks](webappapis.html#concept-task){#worker-event-loop:concept-task}.
These [worker event
loops](webappapis.html#worker-event-loop-2){#worker-event-loop:worker-event-loop-2-2}
are created by the [run a
worker](#run-a-worker){#worker-event-loop:run-a-worker} algorithm.

Each
[`WorkerGlobalScope`{#worker-event-loop:workerglobalscope}](#workerglobalscope)
object has a [closing]{#dom-workerglobalscope-closing .dfn
dfn-for="WorkerGlobalScope" export=""} flag, which must be initially
false, but which can get set to true by the algorithms in the processing
model section below.

Once the
[`WorkerGlobalScope`{#worker-event-loop:workerglobalscope-2}](#workerglobalscope)\'s
[closing](#dom-workerglobalscope-closing){#worker-event-loop:dom-workerglobalscope-closing}
flag is set to true, the [event
loop](webappapis.html#event-loop){#worker-event-loop:event-loop}\'s
[task
queues](webappapis.html#task-queue){#worker-event-loop:task-queue-2}
must discard any further
[tasks](webappapis.html#concept-task){#worker-event-loop:concept-task-2}
that would be added to them (tasks already on the queue are unaffected
except where otherwise specified). Effectively, once the
[closing](#dom-workerglobalscope-closing){#worker-event-loop:dom-workerglobalscope-closing-2}
flag is true, timers stop firing, notifications for all pending
background operations are dropped, etc.

#### [10.2.3]{.secno} The worker\'s lifetime[](#the-worker's-lifetime){.self-link} {#the-worker's-lifetime}

Workers communicate with other workers and with
[`Window`{#the-worker's-lifetime:window}](nav-history-apis.html#window)s
through [message
channels](web-messaging.html#channel-messaging){#the-worker's-lifetime:channel-messaging}
and their
[`MessagePort`{#the-worker's-lifetime:messageport}](web-messaging.html#messageport)
objects.

Each
[`WorkerGlobalScope`{#the-worker's-lifetime:workerglobalscope}](#workerglobalscope)
object `worker global scope`{.variable} has a list of [the worker\'s
ports]{#the-worker's-ports .dfn export=""}, which consists of all the
[`MessagePort`{#the-worker's-lifetime:messageport-2}](web-messaging.html#messageport)
objects that are entangled with another port and that have one (but only
one) port owned by `worker global scope`{.variable}. This list includes
the implicit
[`MessagePort`{#the-worker's-lifetime:messageport-3}](web-messaging.html#messageport)
in the case of [dedicated
workers](#dedicatedworkerglobalscope){#the-worker's-lifetime:dedicatedworkerglobalscope}.

Given an [environment settings
object](webappapis.html#environment-settings-object){#the-worker's-lifetime:environment-settings-object}
`o`{.variable} when creating or obtaining a worker, the [relevant owner
to add]{#relevant-owner-to-add .dfn} depends on the type of [global
object](webappapis.html#concept-settings-object-global){#the-worker's-lifetime:concept-settings-object-global}
specified by `o`{.variable}. If `o`{.variable}\'s [global
object](webappapis.html#concept-settings-object-global){#the-worker's-lifetime:concept-settings-object-global-2}
is a
[`WorkerGlobalScope`{#the-worker's-lifetime:workerglobalscope-2}](#workerglobalscope)
object (i.e., if we are creating a nested dedicated worker), then the
relevant owner is that global object. Otherwise, `o`{.variable}\'s
[global
object](webappapis.html#concept-settings-object-global){#the-worker's-lifetime:concept-settings-object-global-3}
is a
[`Window`{#the-worker's-lifetime:window-2}](nav-history-apis.html#window)
object, and the relevant owner is that
[`Window`{#the-worker's-lifetime:window-3}](nav-history-apis.html#window)\'s
[associated
`Document`](nav-history-apis.html#concept-document-window){#the-worker's-lifetime:concept-document-window}.

------------------------------------------------------------------------

A worker is said to be a [permissible worker]{#permissible-worker .dfn}
if its
[`WorkerGlobalScope`{#the-worker's-lifetime:workerglobalscope-3}](#workerglobalscope)\'s
[owner
set](#concept-WorkerGlobalScope-owner-set){#the-worker's-lifetime:concept-WorkerGlobalScope-owner-set}
is not
[empty](https://infra.spec.whatwg.org/#list-is-empty){#the-worker's-lifetime:list-is-empty
x-internal="list-is-empty"} or:

- its [owner
  set](#concept-WorkerGlobalScope-owner-set){#the-worker's-lifetime:concept-WorkerGlobalScope-owner-set-2}
  has been
  [empty](https://infra.spec.whatwg.org/#list-is-empty){#the-worker's-lifetime:list-is-empty-2
  x-internal="list-is-empty"} for no more than a short
  [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#the-worker's-lifetime:implementation-defined
  x-internal="implementation-defined"} timeout value,
- its
  [`WorkerGlobalScope`{#the-worker's-lifetime:workerglobalscope-4}](#workerglobalscope)
  object is a
  [`SharedWorkerGlobalScope`{#the-worker's-lifetime:sharedworkerglobalscope}](#sharedworkerglobalscope)
  object (i.e., the worker is a shared worker), and
- the user agent has a
  [navigable](document-sequences.html#navigable){#the-worker's-lifetime:navigable}
  whose [active
  document](document-sequences.html#nav-document){#the-worker's-lifetime:nav-document}
  is not [completely
  loaded](document-lifecycle.html#completely-loaded){#the-worker's-lifetime:completely-loaded}.

The second part of this definition allows a shared worker to survive for
a short time while a page is loading, in case that page is going to
contact the shared worker again. This can be used by user agents as a
way to avoid the cost of restarting a shared worker used by a site when
the user is navigating from page to page within that site.

A worker is said to be an [active needed worker]{#active-needed-worker
.dfn} if any of its
[owners](#concept-WorkerGlobalScope-owner-set){#the-worker's-lifetime:concept-WorkerGlobalScope-owner-set-3}
are either
[`Document`{#the-worker's-lifetime:document}](dom.html#document) objects
that are [fully
active](document-sequences.html#fully-active){#the-worker's-lifetime:fully-active}
or [active needed
workers](#active-needed-worker){#the-worker's-lifetime:active-needed-worker}.

A worker is said to be a [protected worker]{#protected-worker .dfn} if
it is an [active needed
worker](#active-needed-worker){#the-worker's-lifetime:active-needed-worker-2}
and either it has outstanding timers, database transactions, or network
connections, or its list of [the worker\'s
ports](#the-worker's-ports){#the-worker's-lifetime:the-worker's-ports}
is not empty, or its
[`WorkerGlobalScope`{#the-worker's-lifetime:workerglobalscope-5}](#workerglobalscope)
is actually a
[`SharedWorkerGlobalScope`{#the-worker's-lifetime:sharedworkerglobalscope-2}](#sharedworkerglobalscope)
object (i.e., the worker is a shared worker).

A worker is said to be a [suspendable worker]{#suspendable-worker .dfn}
if it is not an [active needed
worker](#active-needed-worker){#the-worker's-lifetime:active-needed-worker-3}
but it is a [permissible
worker](#permissible-worker){#the-worker's-lifetime:permissible-worker}.

#### [10.2.4]{.secno} []{#processing-model-10}Processing model[](#worker-processing-model){.self-link} {#worker-processing-model}

When a user agent is to [run a worker]{#run-a-worker .dfn export=""} for
a script with [`Worker`{#worker-processing-model:worker}](#worker) or
[`SharedWorker`{#worker-processing-model:sharedworker}](#sharedworker)
object `worker`{.variable},
[URL](https://url.spec.whatwg.org/#concept-url){#worker-processing-model:url
x-internal="url"} `url`{.variable}, [environment settings
object](webappapis.html#environment-settings-object){#worker-processing-model:environment-settings-object}
`outside settings`{.variable},
[`MessagePort`{#worker-processing-model:messageport}](web-messaging.html#messageport)
`outside port`{.variable}, and a
[`WorkerOptions`{#worker-processing-model:workeroptions}](#workeroptions)
dictionary `options`{.variable}, it must run the following steps.

1.  Let `is shared`{.variable} be true if `worker`{.variable} is a
    [`SharedWorker`{#worker-processing-model:sharedworker-2}](#sharedworker)
    object, and false otherwise.

2.  Let `owner`{.variable} be the [relevant owner to
    add](#relevant-owner-to-add){#worker-processing-model:relevant-owner-to-add}
    given `outside settings`{.variable}.

3.  Let `unsafeWorkerCreationTime`{.variable} be the [unsafe shared
    current
    time](https://w3c.github.io/hr-time/#dfn-unsafe-shared-current-time){#worker-processing-model:unsafe-shared-current-time
    x-internal="unsafe-shared-current-time"}.

4.  Let `agent`{.variable} be the result of [obtaining a
    dedicated/shared worker
    agent](webappapis.html#obtain-a-dedicated/shared-worker-agent){#worker-processing-model:obtain-a-dedicated/shared-worker-agent}
    given `outside settings`{.variable} and `is shared`{.variable}. Run
    the rest of these steps in that agent.

5.  Let `realm execution context`{.variable} be the result of [creating
    a new
    realm](webappapis.html#creating-a-new-javascript-realm){#worker-processing-model:creating-a-new-javascript-realm}
    given `agent`{.variable} and the following customizations:

    - For the global object, if `is shared`{.variable} is true, create a
      new
      [`SharedWorkerGlobalScope`{#worker-processing-model:sharedworkerglobalscope}](#sharedworkerglobalscope)
      object. Otherwise, create a new
      [`DedicatedWorkerGlobalScope`{#worker-processing-model:dedicatedworkerglobalscope}](#dedicatedworkerglobalscope)
      object.

6.  Let `worker global scope`{.variable} be the [global
    object](webappapis.html#concept-realm-global){#worker-processing-model:concept-realm-global}
    of `realm execution context`{.variable}\'s Realm component.

    This is the
    [`DedicatedWorkerGlobalScope`{#worker-processing-model:dedicatedworkerglobalscope-2}](#dedicatedworkerglobalscope)
    or
    [`SharedWorkerGlobalScope`{#worker-processing-model:sharedworkerglobalscope-2}](#sharedworkerglobalscope)
    object created in the previous step.

7.  [Set up a worker environment settings
    object](#set-up-a-worker-environment-settings-object){#worker-processing-model:set-up-a-worker-environment-settings-object}
    with `realm execution context`{.variable},
    `outside settings`{.variable}, and
    `unsafeWorkerCreationTime`{.variable}, and let
    `inside settings`{.variable} be the result.

8.  Set `worker global scope`{.variable}\'s
    [name](#concept-workerglobalscope-name){#worker-processing-model:concept-workerglobalscope-name}
    to the value of `options`{.variable}\'s `name` member.

9.  [Append](https://infra.spec.whatwg.org/#set-append){#worker-processing-model:set-append
    x-internal="set-append"} `owner`{.variable} to
    `worker global scope`{.variable}\'s [owner
    set](#concept-WorkerGlobalScope-owner-set){#worker-processing-model:concept-WorkerGlobalScope-owner-set}.

10. If `is shared`{.variable} is true, then:

    1.  Set `worker global scope`{.variable}\'s [constructor
        origin](#concept-sharedworkerglobalscope-constructor-origin){#worker-processing-model:concept-sharedworkerglobalscope-constructor-origin}
        to `outside settings`{.variable}\'s
        [origin](webappapis.html#concept-settings-object-origin){#worker-processing-model:concept-settings-object-origin}.

    2.  Set `worker global scope`{.variable}\'s [constructor
        url](#concept-sharedworkerglobalscope-constructor-url){#worker-processing-model:concept-sharedworkerglobalscope-constructor-url}
        to `url`{.variable}.

    3.  Set `worker global scope`{.variable}\'s
        [type](#concept-workerglobalscope-type){#worker-processing-model:concept-workerglobalscope-type}
        to the value of `options`{.variable}\'s `type` member.

    4.  Set `worker global scope`{.variable}\'s
        [credentials](#concept-sharedworkerglobalscope-credentials){#worker-processing-model:concept-sharedworkerglobalscope-credentials}
        to the value of `options`{.variable}\'s `credentials` member.

11. Let `destination`{.variable} be \"`sharedworker`\" if
    `is shared`{.variable} is true, and \"`worker`\" otherwise.

12. Obtain `script`{.variable} by switching on the value of
    `options`{.variable}\'s `type` member:

    \"`classic`\"
    :   [Fetch a classic worker
        script](webappapis.html#fetch-a-classic-worker-script){#worker-processing-model:fetch-a-classic-worker-script}
        given `url`{.variable}, `outside settings`{.variable},
        `destination`{.variable}, `inside settings`{.variable}, and with
        `onComplete`{.variable} and `performFetch`{.variable} as defined
        below.

    \"`module`\"
    :   [Fetch a module worker script
        graph](webappapis.html#fetch-a-module-worker-script-tree){#worker-processing-model:fetch-a-module-worker-script-tree}
        given `url`{.variable}, `outside settings`{.variable},
        `destination`{.variable}, the value of the `credentials` member
        of `options`{.variable}, `inside settings`{.variable}, and with
        `onComplete`{.variable} and `performFetch`{.variable} as defined
        below.

    In both cases, let `performFetch`{.variable} be the following
    [perform the fetch
    hook](webappapis.html#fetching-scripts-perform-fetch){#worker-processing-model:fetching-scripts-perform-fetch}
    given `request`{.variable},
    *[isTopLevel](webappapis.html#fetching-scripts-is-top-level)* and
    *[processCustomFetchResponse](webappapis.html#fetching-scripts-processcustomfetchresponse)*:

    1.  If `isTopLevel`{.variable} is false,
        [fetch](https://fetch.spec.whatwg.org/#concept-fetch){#worker-processing-model:concept-fetch
        x-internal="concept-fetch"} `request`{.variable} with
        *[processResponseConsumeBody](https://fetch.spec.whatwg.org/#process-response-end-of-body){x-internal="processresponseconsumebody"}*
        set to `processCustomFetchResponse`{.variable}, and abort these
        steps.

    2.  Set `request`{.variable}\'s [reserved
        client](https://fetch.spec.whatwg.org/#concept-request-reserved-client){#worker-processing-model:concept-request-reserved-client
        x-internal="concept-request-reserved-client"} to
        `inside settings`{.variable}.

    3.  [Fetch](https://fetch.spec.whatwg.org/#concept-fetch){#worker-processing-model:concept-fetch-2
        x-internal="concept-fetch"} `request`{.variable} with
        *[processResponseConsumeBody](https://fetch.spec.whatwg.org/#process-response-end-of-body){x-internal="processresponseconsumebody"}*
        set to the following steps given
        [response](https://fetch.spec.whatwg.org/#concept-response){#worker-processing-model:concept-response
        x-internal="concept-response"} `response`{.variable} and null,
        failure, or a [byte
        sequence](https://infra.spec.whatwg.org/#byte-sequence){#worker-processing-model:byte-sequence
        x-internal="byte-sequence"} `bodyBytes`{.variable}:

        1.  Set `worker global scope`{.variable}\'s
            [url](#concept-workerglobalscope-url){#worker-processing-model:concept-workerglobalscope-url}
            to `response`{.variable}\'s
            [url](https://fetch.spec.whatwg.org/#concept-response-url){#worker-processing-model:concept-response-url
            x-internal="concept-response-url"}.

        2.  [Initialize worker global scope\'s policy
            container](browsers.html#initialize-worker-policy-container){#worker-processing-model:initialize-worker-policy-container}
            given `worker global scope`{.variable},
            `response`{.variable}, and `inside settings`{.variable}.

        3.  If the [Run CSP initialization for a global
            object](https://w3c.github.io/webappsec-csp/#run-global-object-csp-initialization){#worker-processing-model:run-csp-initialization-for-a-global-object
            x-internal="run-csp-initialization-for-a-global-object"}
            algorithm returns \"`Blocked`\" when executed upon
            `worker global scope`{.variable}, set `response`{.variable}
            to a [network
            error](https://fetch.spec.whatwg.org/#concept-network-error){#worker-processing-model:network-error
            x-internal="network-error"}.
            [\[CSP\]](references.html#refsCSP)

        4.  If `worker global scope`{.variable}\'s [embedder
            policy](#concept-workerglobalscope-embedder-policy){#worker-processing-model:concept-workerglobalscope-embedder-policy}\'s
            [value](browsers.html#embedder-policy-value-2){#worker-processing-model:embedder-policy-value-2}
            is [compatible with cross-origin
            isolation](browsers.html#compatible-with-cross-origin-isolation){#worker-processing-model:compatible-with-cross-origin-isolation}
            and `is shared`{.variable} is true, then set
            `agent`{.variable}\'s [agent
            cluster](https://tc39.es/ecma262/#sec-agent-clusters){#worker-processing-model:agent-cluster
            x-internal="agent-cluster"}\'s [cross-origin isolation
            mode](webappapis.html#agent-cluster-cross-origin-isolation){#worker-processing-model:agent-cluster-cross-origin-isolation}
            to
            \"[`logical`{#worker-processing-model:cross-origin-isolation-logical}](document-sequences.html#cross-origin-isolation-logical)\"
            or
            \"[`concrete`{#worker-processing-model:cross-origin-isolation-concrete}](document-sequences.html#cross-origin-isolation-concrete)\".
            The one chosen is
            [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#worker-processing-model:implementation-defined
            x-internal="implementation-defined"}.

            This really ought to be set when the agent cluster is
            created, which requires a redesign of this section.

        5.  If the result of [checking a global object\'s embedder
            policy](browsers.html#check-a-global-object's-embedder-policy){#worker-processing-model:check-a-global-object's-embedder-policy}
            with `worker global scope`{.variable},
            `outside settings`{.variable}, and `response`{.variable} is
            false, then set `response`{.variable} to a [network
            error](https://fetch.spec.whatwg.org/#concept-network-error){#worker-processing-model:network-error-2
            x-internal="network-error"}.

        6.  Set `worker global scope`{.variable}\'s [cross-origin
            isolated
            capability](#concept-workerglobalscope-cross-origin-isolated-capability){#worker-processing-model:concept-workerglobalscope-cross-origin-isolated-capability}
            to true if `agent`{.variable}\'s [agent
            cluster](https://tc39.es/ecma262/#sec-agent-clusters){#worker-processing-model:agent-cluster-2
            x-internal="agent-cluster"}\'s [cross-origin isolation
            mode](webappapis.html#agent-cluster-cross-origin-isolation){#worker-processing-model:agent-cluster-cross-origin-isolation-2}
            is
            \"[`concrete`{#worker-processing-model:cross-origin-isolation-concrete-2}](document-sequences.html#cross-origin-isolation-concrete)\".

        7.  If `is shared`{.variable} is false and `owner`{.variable}\'s
            [cross-origin isolated
            capability](webappapis.html#concept-settings-object-cross-origin-isolated-capability){#worker-processing-model:concept-settings-object-cross-origin-isolated-capability}
            is false, then set `worker global scope`{.variable}\'s
            [cross-origin isolated
            capability](#concept-workerglobalscope-cross-origin-isolated-capability){#worker-processing-model:concept-workerglobalscope-cross-origin-isolated-capability-2}
            to false.

        8.  If `is shared`{.variable} is false and
            `response`{.variable}\'s
            [url](https://fetch.spec.whatwg.org/#concept-response-url){#worker-processing-model:concept-response-url-2
            x-internal="concept-response-url"}\'s
            [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#worker-processing-model:concept-url-scheme
            x-internal="concept-url-scheme"} is \"`data`\", then set
            `worker global scope`{.variable}\'s [cross-origin isolated
            capability](#concept-workerglobalscope-cross-origin-isolated-capability){#worker-processing-model:concept-workerglobalscope-cross-origin-isolated-capability-3}
            to false.

            This is a conservative default for now, while we figure out
            how workers in general, and
            [`data:`{#worker-processing-model:data-protocol}](https://www.rfc-editor.org/rfc/rfc2397#section-2){x-internal="data-protocol"}
            URL workers in particular (which are cross-origin from their
            owner), will be treated in the context of permissions
            policies. See [w3c/webappsec-permissions-policy issue
            #207](https://github.com/w3c/webappsec-permissions-policy/issues/207)
            for more details.

        9.  Run `processCustomFetchResponse`{.variable} with
            `response`{.variable} and `bodyBytes`{.variable}.

    In both cases, let `onComplete`{.variable} given `script`{.variable}
    be the following steps:

    1.  If `script`{.variable} is null or if `script`{.variable}\'s
        [error to
        rethrow](webappapis.html#concept-script-error-to-rethrow){#worker-processing-model:concept-script-error-to-rethrow}
        is non-null, then:

        1.  [Queue a global
            task](webappapis.html#queue-a-global-task){#worker-processing-model:queue-a-global-task}
            on the [DOM manipulation task
            source](webappapis.html#dom-manipulation-task-source){#worker-processing-model:dom-manipulation-task-source}
            given `worker`{.variable}\'s [relevant global
            object](webappapis.html#concept-relevant-global){#worker-processing-model:concept-relevant-global}
            to [fire an
            event](https://dom.spec.whatwg.org/#concept-event-fire){#worker-processing-model:concept-event-fire
            x-internal="concept-event-fire"} named
            [`error`{#worker-processing-model:event-error}](indices.html#event-error)
            at `worker`{.variable}.

        2.  Run the [environment discarding
            steps](webappapis.html#environment-discarding-steps){#worker-processing-model:environment-discarding-steps}
            for `inside settings`{.variable}.

        3.  Abort these steps.

    2.  Associate `worker`{.variable} with
        `worker global scope`{.variable}.

    3.  Let `inside port`{.variable} be a
        [new](https://webidl.spec.whatwg.org/#new){#worker-processing-model:new
        x-internal="new"}
        [`MessagePort`{#worker-processing-model:messageport-2}](web-messaging.html#messageport)
        object in `inside settings`{.variable}\'s
        [realm](webappapis.html#environment-settings-object's-realm){#worker-processing-model:environment-settings-object's-realm}.

    4.  If `is shared`{.variable} is false, then:

        1.  Set `inside port`{.variable}\'s [message event
            target](web-messaging.html#message-event-target){#worker-processing-model:message-event-target}
            to `worker global scope`{.variable}.

        2.  Set `worker global scope`{.variable}\'s [inside
            port](#inside-port){#worker-processing-model:inside-port} to
            `inside port`{.variable}.

    5.  [Entangle](web-messaging.html#entangle){#worker-processing-model:entangle}
        `outside port`{.variable} and `inside port`{.variable}.

    6.  Create a new
        [`WorkerLocation`{#worker-processing-model:workerlocation}](#workerlocation)
        object and associate it with `worker global scope`{.variable}.

    7.  **Closing orphan workers**: Start monitoring the worker such
        that no sooner than it stops being a [protected
        worker](#protected-worker){#worker-processing-model:protected-worker},
        and no later than it stops being a [permissible
        worker](#permissible-worker){#worker-processing-model:permissible-worker},
        `worker global scope`{.variable}\'s
        [closing](#dom-workerglobalscope-closing){#worker-processing-model:dom-workerglobalscope-closing}
        flag is set to true.

    8.  **Suspending workers**: Start monitoring the worker, such that
        whenever `worker global scope`{.variable}\'s
        [closing](#dom-workerglobalscope-closing){#worker-processing-model:dom-workerglobalscope-closing-2}
        flag is false and the worker is a [suspendable
        worker](#suspendable-worker){#worker-processing-model:suspendable-worker},
        the user agent suspends execution of script in that worker until
        such time as either the
        [closing](#dom-workerglobalscope-closing){#worker-processing-model:dom-workerglobalscope-closing-3}
        flag switches to true or the worker stops being a [suspendable
        worker](#suspendable-worker){#worker-processing-model:suspendable-worker-2}.

    9.  Set `inside settings`{.variable}\'s [execution ready
        flag](webappapis.html#concept-environment-execution-ready-flag){#worker-processing-model:concept-environment-execution-ready-flag}.

    10. If `script`{.variable} is a [classic
        script](webappapis.html#classic-script){#worker-processing-model:classic-script},
        then [run the classic
        script](webappapis.html#run-a-classic-script){#worker-processing-model:run-a-classic-script}
        `script`{.variable}. Otherwise, it is a [module
        script](webappapis.html#module-script){#worker-processing-model:module-script};
        [run the module
        script](webappapis.html#run-a-module-script){#worker-processing-model:run-a-module-script}
        `script`{.variable}.

        In addition to the usual possibilities of returning a value or
        failing due to an exception, this could be [prematurely
        aborted](webappapis.html#abort-a-running-script){#worker-processing-model:abort-a-running-script}
        by the [terminate a
        worker](#terminate-a-worker){#worker-processing-model:terminate-a-worker}
        algorithm defined below.

    11. Enable `outside port`{.variable}\'s [port message
        queue](web-messaging.html#port-message-queue){#worker-processing-model:port-message-queue}.

    12. If `is shared`{.variable} is false, enable the [port message
        queue](web-messaging.html#port-message-queue){#worker-processing-model:port-message-queue-2}
        of the worker\'s implicit port.

    13. If `is shared`{.variable} is true, then [queue a global
        task](webappapis.html#queue-a-global-task){#worker-processing-model:queue-a-global-task-2}
        on [DOM manipulation task
        source](webappapis.html#dom-manipulation-task-source){#worker-processing-model:dom-manipulation-task-source-2}
        given `worker global scope`{.variable} to [fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#worker-processing-model:concept-event-fire-2
        x-internal="concept-event-fire"} named
        [`connect`{#worker-processing-model:event-workerglobalscope-connect}](indices.html#event-workerglobalscope-connect)
        at `worker global scope`{.variable}, using
        [`MessageEvent`{#worker-processing-model:messageevent}](comms.html#messageevent),
        with the
        [`data`{#worker-processing-model:dom-messageevent-data}](comms.html#dom-messageevent-data)
        attribute initialized to the empty string, the
        [`ports`{#worker-processing-model:dom-messageevent-ports}](comms.html#dom-messageevent-ports)
        attribute initialized to a new [frozen
        array](https://webidl.spec.whatwg.org/#dfn-frozen-array-type){#worker-processing-model:frozen-array
        x-internal="frozen-array"} containing `inside port`{.variable},
        and the
        [`source`{#worker-processing-model:dom-messageevent-source}](comms.html#dom-messageevent-source)
        attribute initialized to `inside port`{.variable}.

    14. Enable the [client message
        queue](https://w3c.github.io/ServiceWorker/#dfn-client-message-queue){#worker-processing-model:dfn-client-message-queue
        x-internal="dfn-client-message-queue"} of the
        [`ServiceWorkerContainer`{#worker-processing-model:serviceworkercontainer}](https://w3c.github.io/ServiceWorker/#serviceworkercontainer){x-internal="serviceworkercontainer"}
        object whose associated [service worker
        client](https://w3c.github.io/ServiceWorker/#serviceworkercontainer-service-worker-client){#worker-processing-model:serviceworkercontainer-service-worker-client
        x-internal="serviceworkercontainer-service-worker-client"} is
        `worker global scope`{.variable}\'s [relevant settings
        object](webappapis.html#relevant-settings-object){#worker-processing-model:relevant-settings-object}.

    15. **Event loop**: Run the [responsible event
        loop](webappapis.html#responsible-event-loop){#worker-processing-model:responsible-event-loop}
        specified by `inside settings`{.variable} until it is destroyed.

        The handling of events or the execution of callbacks by
        [tasks](webappapis.html#concept-task){#worker-processing-model:concept-task}
        run by the [event
        loop](webappapis.html#event-loop){#worker-processing-model:event-loop}
        might get [prematurely
        aborted](webappapis.html#abort-a-running-script){#worker-processing-model:abort-a-running-script-2}
        by the [terminate a
        worker](#terminate-a-worker){#worker-processing-model:terminate-a-worker-2}
        algorithm defined below.

        The worker processing model remains on this step until the event
        loop is destroyed, which happens after the
        [closing](#dom-workerglobalscope-closing){#worker-processing-model:dom-workerglobalscope-closing-4}
        flag is set to true, as described in the [event
        loop](webappapis.html#event-loop){#worker-processing-model:event-loop-2}
        processing model.

    16. [Clear](https://infra.spec.whatwg.org/#map-clear){#worker-processing-model:map-clear
        x-internal="map-clear"} the `worker global scope`{.variable}\'s
        [map of active
        timers](timers-and-user-prompts.html#map-of-active-timers){#worker-processing-model:map-of-active-timers}.

    17. Disentangle all the ports in the list of [the worker\'s
        ports](#the-worker's-ports){#worker-processing-model:the-worker's-ports}.

    18. [Empty](https://infra.spec.whatwg.org/#list-empty){#worker-processing-model:list-empty
        x-internal="list-empty"} `worker global scope`{.variable}\'s
        [owner
        set](#concept-WorkerGlobalScope-owner-set){#worker-processing-model:concept-WorkerGlobalScope-owner-set-2}.

------------------------------------------------------------------------

When a user agent is to [terminate a worker]{#terminate-a-worker .dfn
export=""}, it must run the following steps [in
parallel](infrastructure.html#in-parallel){#worker-processing-model:in-parallel}
with the worker\'s main loop (the \"[run a
worker](#run-a-worker){#worker-processing-model:run-a-worker}\"
processing model defined above):

1.  Set the worker\'s
    [`WorkerGlobalScope`{#worker-processing-model:workerglobalscope}](#workerglobalscope)
    object\'s
    [closing](#dom-workerglobalscope-closing){#worker-processing-model:dom-workerglobalscope-closing-5}
    flag to true.

2.  If there are any
    [tasks](webappapis.html#concept-task){#worker-processing-model:concept-task-2}
    queued in the
    [`WorkerGlobalScope`{#worker-processing-model:workerglobalscope-2}](#workerglobalscope)
    object\'s [relevant
    agent](webappapis.html#relevant-agent){#worker-processing-model:relevant-agent}\'s
    [event
    loop](webappapis.html#concept-agent-event-loop){#worker-processing-model:concept-agent-event-loop}\'s
    [task
    queues](webappapis.html#task-queue){#worker-processing-model:task-queue},
    discard them without processing them.

3.  [Abort the
    script](webappapis.html#abort-a-running-script){#worker-processing-model:abort-a-running-script-3}
    currently running in the worker.

4.  If the worker\'s
    [`WorkerGlobalScope`{#worker-processing-model:workerglobalscope-3}](#workerglobalscope)
    object is actually a
    [`DedicatedWorkerGlobalScope`{#worker-processing-model:dedicatedworkerglobalscope-3}](#dedicatedworkerglobalscope)
    object (i.e. the worker is a dedicated worker), then empty the [port
    message
    queue](web-messaging.html#port-message-queue){#worker-processing-model:port-message-queue-3}
    of the port that the worker\'s implicit port is entangled with.

User agents may invoke the [terminate a
worker](#terminate-a-worker){#worker-processing-model:terminate-a-worker-3}
algorithm when a worker stops being an [active needed
worker](#active-needed-worker){#worker-processing-model:active-needed-worker}
and the worker continues executing even after its
[closing](#dom-workerglobalscope-closing){#worker-processing-model:dom-workerglobalscope-closing-6}
flag was set to true.

#### [10.2.5]{.secno} Runtime script errors[](#runtime-script-errors-2){.self-link} {#runtime-script-errors-2}

Whenever an uncaught runtime script error occurs in one of the worker\'s
scripts, if the error did not occur while handling a previous script
error, the user agent will
[report](webappapis.html#report-an-exception){#runtime-script-errors-2:report-an-exception}
it for the worker\'s
[`WorkerGlobalScope`{#runtime-script-errors-2:workerglobalscope}](#workerglobalscope)
object.

#### [10.2.6]{.secno} Creating workers[](#creating-workers){.self-link}

##### [10.2.6.1]{.secno} []{#the-abstractworker-abstract-interface}The [`AbstractWorker`{#the-abstractworker-mixin:abstractworker}](#abstractworker) mixin[](#the-abstractworker-mixin){.self-link}

``` idl
interface mixin AbstractWorker {
  attribute EventHandler onerror;
};
```

The following are the [event
handlers](webappapis.html#event-handlers){#the-abstractworker-mixin:event-handlers}
(and their corresponding [event handler event
types](webappapis.html#event-handler-event-type){#the-abstractworker-mixin:event-handler-event-type})
that must be supported, as [event handler IDL
attributes](webappapis.html#event-handler-idl-attributes){#the-abstractworker-mixin:event-handler-idl-attributes},
by objects implementing the
[`AbstractWorker`{#the-abstractworker-mixin:abstractworker-2}](#abstractworker)
interface:

[Event
handler](webappapis.html#event-handlers){#the-abstractworker-mixin:event-handlers-2}

[Event handler event
type](webappapis.html#event-handler-event-type){#the-abstractworker-mixin:event-handler-event-type-2}

[`onerror`]{#handler-abstractworker-onerror .dfn
dfn-for="AbstractWorker" dfn-type="attribute"}

::::::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[ServiceWorker/error_event](https://developer.mozilla.org/en-US/docs/Web/API/ServiceWorker/error_event "The error event fires whenever an error occurs in the service worker.")

Support in all current engines.

::: support
[Firefox44+]{.firefox .yes}[Safari11.1+]{.safari
.yes}[Chrome40+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)17+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::

:::: feature
[SharedWorker/error_event](https://developer.mozilla.org/en-US/docs/Web/API/SharedWorker/error_event "The error event of the SharedWorker interface fires when an error occurs in the worker.")

Support in all current engines.

::: support
[Firefox29+]{.firefox .yes}[Safari16+]{.safari .yes}[Chrome5+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android33+]{.firefox_android .yes}[Safari iOS16+]{.safari_ios
.yes}[Chrome AndroidNo]{.chrome_android .no}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet4.0--5.0]{.samsunginternet_android .no}[Opera
Android11--14]{.opera_android .no}
:::
::::

:::: feature
[Worker/error_event](https://developer.mozilla.org/en-US/docs/Web/API/Worker/error_event "The error event of the Worker interface fires when an error occurs in the worker.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::::::

[`error`{#the-abstractworker-mixin:event-error}](indices.html#event-error)

##### [10.2.6.2]{.secno} Script settings for workers[](#script-settings-for-workers){.self-link}

To [set up a worker environment settings
object]{#set-up-a-worker-environment-settings-object .dfn}, given a
[JavaScript execution
context](https://tc39.es/ecma262/#sec-execution-contexts){#script-settings-for-workers:javascript-execution-context
x-internal="javascript-execution-context"}
`execution context`{.variable}, an [environment settings
object](webappapis.html#environment-settings-object){#script-settings-for-workers:environment-settings-object}
`outside settings`{.variable}, and a number
`unsafeWorkerCreationTime`{.variable}:

1.  Let `inherited origin`{.variable} be
    `outside settings`{.variable}\'s
    [origin](webappapis.html#concept-settings-object-origin){#script-settings-for-workers:concept-settings-object-origin}.

2.  Let `realm`{.variable} be the value of
    `execution context`{.variable}\'s Realm component.

3.  Let `worker global scope`{.variable} be `realm`{.variable}\'s
    [global
    object](webappapis.html#concept-realm-global){#script-settings-for-workers:concept-realm-global}.

4.  Let `settings object`{.variable} be a new [environment settings
    object](webappapis.html#environment-settings-object){#script-settings-for-workers:environment-settings-object-2}
    whose algorithms are defined as follows:

    The [realm execution context](webappapis.html#realm-execution-context){#script-settings-for-workers:realm-execution-context}
    :   Return `execution context`{.variable}.

    The [module map](webappapis.html#concept-settings-object-module-map){#script-settings-for-workers:concept-settings-object-module-map}

    :   Return `worker global scope`{.variable}\'s [module
        map](#concept-workerglobalscope-module-map){#script-settings-for-workers:concept-workerglobalscope-module-map}.

    The [API base URL](webappapis.html#api-base-url){#script-settings-for-workers:api-base-url}

    :   Return `worker global scope`{.variable}\'s
        [url](#concept-workerglobalscope-url){#script-settings-for-workers:concept-workerglobalscope-url}.

    The [origin](webappapis.html#concept-settings-object-origin){#script-settings-for-workers:concept-settings-object-origin-2}

    :   Return a unique [opaque
        origin](browsers.html#concept-origin-opaque){#script-settings-for-workers:concept-origin-opaque}
        if `worker global scope`{.variable}\'s
        [url](#concept-workerglobalscope-url){#script-settings-for-workers:concept-workerglobalscope-url-2}\'s
        [scheme](https://url.spec.whatwg.org/#concept-url-scheme){#script-settings-for-workers:concept-url-scheme
        x-internal="concept-url-scheme"} is \"`data`\", and
        `inherited origin`{.variable} otherwise.

    The [policy container](webappapis.html#concept-settings-object-policy-container){#script-settings-for-workers:concept-settings-object-policy-container}

    :   Return `worker global scope`{.variable}\'s [policy
        container](#concept-workerglobalscope-policy-container){#script-settings-for-workers:concept-workerglobalscope-policy-container}.

    The [cross-origin isolated capability](webappapis.html#concept-settings-object-cross-origin-isolated-capability){#script-settings-for-workers:concept-settings-object-cross-origin-isolated-capability}
    :   Return `worker global scope`{.variable}\'s [cross-origin
        isolated
        capability](#concept-workerglobalscope-cross-origin-isolated-capability){#script-settings-for-workers:concept-workerglobalscope-cross-origin-isolated-capability}.

    The [time origin](webappapis.html#concept-settings-object-time-origin){#script-settings-for-workers:concept-settings-object-time-origin}

    :   Return the result of
        [coarsening](https://w3c.github.io/hr-time/#dfn-coarsen-time){#script-settings-for-workers:coarsen-time
        x-internal="coarsen-time"} `unsafeWorkerCreationTime`{.variable}
        with `worker global scope`{.variable}\'s [cross-origin isolated
        capability](#concept-workerglobalscope-cross-origin-isolated-capability){#script-settings-for-workers:concept-workerglobalscope-cross-origin-isolated-capability-2}.

5.  Set `settings object`{.variable}\'s
    [id](webappapis.html#concept-environment-id){#script-settings-for-workers:concept-environment-id}
    to a new unique opaque string, [creation
    URL](webappapis.html#concept-environment-creation-url){#script-settings-for-workers:concept-environment-creation-url}
    to `worker global scope`{.variable}\'s
    [url](https://url.spec.whatwg.org/#concept-url){#script-settings-for-workers:url
    x-internal="url"}, [top-level creation
    URL](webappapis.html#concept-environment-top-level-creation-url){#script-settings-for-workers:concept-environment-top-level-creation-url}
    to null, [target browsing
    context](webappapis.html#concept-environment-target-browsing-context){#script-settings-for-workers:concept-environment-target-browsing-context}
    to null, and [active service
    worker](webappapis.html#concept-environment-active-service-worker){#script-settings-for-workers:concept-environment-active-service-worker}
    to null.

6.  If `worker global scope`{.variable} is a
    [`DedicatedWorkerGlobalScope`{#script-settings-for-workers:dedicatedworkerglobalscope}](#dedicatedworkerglobalscope)
    object, then set `settings object`{.variable}\'s [top-level
    origin](webappapis.html#concept-environment-top-level-origin){#script-settings-for-workers:concept-environment-top-level-origin}
    to `outside settings`{.variable}\'s [top-level
    origin](webappapis.html#concept-environment-top-level-origin){#script-settings-for-workers:concept-environment-top-level-origin-2}.

7.  Otherwise, set `settings object`{.variable}\'s [top-level
    origin](webappapis.html#concept-environment-top-level-origin){#script-settings-for-workers:concept-environment-top-level-origin-3}
    to an
    [implementation-defined](https://infra.spec.whatwg.org/#implementation-defined){#script-settings-for-workers:implementation-defined
    x-internal="implementation-defined"} value.

    See [Client-Side Storage
    Partitioning](https://privacycg.github.io/storage-partitioning/) for
    the latest on properly defining this.

8.  Set `realm`{.variable}\'s \[\[HostDefined\]\] field to
    `settings object`{.variable}.

9.  Return `settings object`{.variable}.

##### [10.2.6.3]{.secno} Dedicated workers and the [`Worker`{#worker-dev}](#worker) interface[](#dedicated-workers-and-the-worker-interface){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[Worker](https://developer.mozilla.org/en-US/docs/Web/API/Worker "The Worker interface of the Web Workers API represents a background task that can be created via script, which can send messages back to its creator.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome2+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

``` idl
[Exposed=(Window,DedicatedWorker,SharedWorker)]
interface Worker : EventTarget {
  constructor((TrustedScriptURL or USVString) scriptURL, optional WorkerOptions options = {});

  undefined terminate();

  undefined postMessage(any message, sequence<object> transfer);
  undefined postMessage(any message, optional StructuredSerializeOptions options = {});
};

dictionary WorkerOptions {
  WorkerType type = "classic";
  RequestCredentials credentials = "same-origin"; // credentials is only used if type is "module"
  DOMString name = "";
};

enum WorkerType { "classic", "module" };

Worker includes AbstractWorker;
Worker includes MessageEventTarget;
```

`worker`{.variable}` = new `[`Worker`](#dom-worker){#dom-worker-dev}`(``scriptURL`{.variable}` [, ``options`{.variable}` ])`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Worker/Worker](https://developer.mozilla.org/en-US/docs/Web/API/Worker/Worker "The Worker() constructor creates a Worker object that executes the script at the specified URL. This script must obey the same-origin policy.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

Returns a new
[`Worker`{#dedicated-workers-and-the-worker-interface:worker-3}](#worker)
object. `scriptURL`{.variable} will be fetched and executed in the
background, creating a new global environment for which
`worker`{.variable} represents the communication channel.
`options`{.variable} can be used to define the
[name](#concept-workerglobalscope-name){#dedicated-workers-and-the-worker-interface:concept-workerglobalscope-name}
of that global environment via the `name` option, primarily for
debugging purposes. It can also ensure this new global environment
supports JavaScript modules (specify `type: "module"`), and if that is
specified, can also be used to specify how `scriptURL`{.variable} is
fetched through the `credentials` option.

`worker`{.variable}`.`[`terminate`](#dom-worker-terminate){#dom-worker-terminate-dev}`()`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Worker/terminate](https://developer.mozilla.org/en-US/docs/Web/API/Worker/terminate "The terminate() method of the Worker interface immediately terminates the Worker. This does not offer the worker an opportunity to finish its operations; it is stopped at once.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome2+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

Aborts `worker`{.variable}\'s associated global environment.

`worker`{.variable}`.`[`postMessage`](#dom-worker-postmessage){#dom-worker-postmessage-dev}`(``message`{.variable}` [, ``transfer`{.variable}` ])`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[Worker/postMessage](https://developer.mozilla.org/en-US/docs/Web/API/Worker/postMessage "The postMessage() method of the Worker interface sends a message to the worker. The first parameter is the data to send to the worker. The data may be any JavaScript object that can be handled by the structured clone algorithm.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome2+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android11+]{.opera_android .yes}
:::
::::
:::::

`worker`{.variable}`.`[`postMessage`](#dom-worker-postmessage-options){#dom-worker-postmessage-options-dev}`(``message`{.variable}` [, { `[`transfer`](web-messaging.html#dom-structuredserializeoptions-transfer){#dedicated-workers-and-the-worker-interface:dom-structuredserializeoptions-transfer}` } ])`

Clones `message`{.variable} and transmits it to `worker`{.variable}\'s
global environment. `transfer`{.variable} can be passed as a list of
objects that are to be transferred rather than cloned.

Each
[`Worker`{#dedicated-workers-and-the-worker-interface:worker-4}](#worker)
object has an associated [outside port]{#outside-port .dfn
dfn-for="Worker"} (a
[`MessagePort`{#dedicated-workers-and-the-worker-interface:messageport}](web-messaging.html#messageport)).
This port is part of a channel that is set up when the worker is
created, but it is not exposed. This object must never be garbage
collected before the
[`Worker`{#dedicated-workers-and-the-worker-interface:worker-5}](#worker)
object.

The [`terminate()`]{#dom-worker-terminate .dfn dfn-for="Worker"
dfn-type="method"} method, when invoked, must cause the [terminate a
worker](#terminate-a-worker){#dedicated-workers-and-the-worker-interface:terminate-a-worker}
algorithm to be run on the worker with which the object is associated.

The
[`postMessage(``message`{.variable}`, ``transfer`{.variable}`)`]{#dom-worker-postmessage
.dfn dfn-for="Worker" dfn-type="method"} and
[`postMessage(``message`{.variable}`, ``options`{.variable}`)`]{#dom-worker-postmessage-options
.dfn dfn-for="Worker" dfn-type="method"} methods on
[`Worker`{#dedicated-workers-and-the-worker-interface:worker-6}](#worker)
objects act as if, when invoked, they immediately invoked the respective
[`postMessage(`{#dedicated-workers-and-the-worker-interface:dom-messageport-postmessage}`message`{.variable}`, `{#dedicated-workers-and-the-worker-interface:dom-messageport-postmessage}`transfer`{.variable}`)`{#dedicated-workers-and-the-worker-interface:dom-messageport-postmessage}](web-messaging.html#dom-messageport-postmessage)
and
[`postMessage(`{#dedicated-workers-and-the-worker-interface:dom-messageport-postmessage-options}`message`{.variable}`, `{#dedicated-workers-and-the-worker-interface:dom-messageport-postmessage-options}`options`{.variable}`)`{#dedicated-workers-and-the-worker-interface:dom-messageport-postmessage-options}](web-messaging.html#dom-messageport-postmessage-options)
on
[this](https://webidl.spec.whatwg.org/#this){#dedicated-workers-and-the-worker-interface:this
x-internal="this"}\'s [outside
port](#outside-port){#dedicated-workers-and-the-worker-interface:outside-port},
with the same arguments, and returned the same return value.

::: example
The
[`postMessage()`{#dedicated-workers-and-the-worker-interface:dom-worker-postmessage-2}](#dom-worker-postmessage)
method\'s first argument can be structured data:

``` js
worker.postMessage({opcode: 'activate', device: 1938, parameters: [23, 102]});
```
:::

------------------------------------------------------------------------

When the
[`Worker(``scriptURL`{.variable}`, ``options`{.variable}`)`]{#dom-worker
.dfn dfn-for="Worker" dfn-type="constructor"} constructor is invoked,
the user agent must run the following steps:

1.  Let `compliantScriptURL`{.variable} be the result of invoking the
    [Get Trusted Type compliant
    string](https://w3c.github.io/trusted-types/dist/spec/#get-trusted-type-compliant-string-algorithm){#dedicated-workers-and-the-worker-interface:tt-getcompliantstring
    x-internal="tt-getcompliantstring"} algorithm with
    [`TrustedScriptURL`{#dedicated-workers-and-the-worker-interface:tt-trustedscripturl-2}](https://w3c.github.io/trusted-types/dist/spec/#trustedscripturl){x-internal="tt-trustedscripturl"},
    [this](https://webidl.spec.whatwg.org/#this){#dedicated-workers-and-the-worker-interface:this-2
    x-internal="this"}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#dedicated-workers-and-the-worker-interface:concept-relevant-global},
    `scriptURL`{.variable}, \"`Worker constructor`\", and \"`script`\".

2.  Let `outside settings`{.variable} be the [current settings
    object](webappapis.html#current-settings-object){#dedicated-workers-and-the-worker-interface:current-settings-object}.

3.  Let `worker URL`{.variable} be the result of [encoding-parsing a
    URL](urls-and-fetching.html#encoding-parsing-a-url){#dedicated-workers-and-the-worker-interface:encoding-parsing-a-url}
    given `compliantScriptURL`{.variable}, relative to
    `outside settings`{.variable}.

    Any
    [same-origin](browsers.html#same-origin){#dedicated-workers-and-the-worker-interface:same-origin}
    URL (including
    [`blob:`{#dedicated-workers-and-the-worker-interface:blob-protocol}](https://w3c.github.io/FileAPI/#DefinitionOfScheme){x-internal="blob-protocol"}
    URLs) can be used.
    [`data:`{#dedicated-workers-and-the-worker-interface:data-protocol}](https://www.rfc-editor.org/rfc/rfc2397#section-2){x-internal="data-protocol"}
    URLs can also be used, but they create a worker with an [opaque
    origin](browsers.html#concept-origin-opaque){#dedicated-workers-and-the-worker-interface:concept-origin-opaque}.

4.  If `worker URL`{.variable} is failure, then throw a
    [\"`SyntaxError`\"](https://webidl.spec.whatwg.org/#syntaxerror){#dedicated-workers-and-the-worker-interface:syntaxerror
    x-internal="syntaxerror"}
    [`DOMException`{#dedicated-workers-and-the-worker-interface:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

5.  Let `worker`{.variable} be a new
    [`Worker`{#dedicated-workers-and-the-worker-interface:worker-7}](#worker)
    object.

6.  Let `outside port`{.variable} be a
    [new](https://webidl.spec.whatwg.org/#new){#dedicated-workers-and-the-worker-interface:new
    x-internal="new"}
    [`MessagePort`{#dedicated-workers-and-the-worker-interface:messageport-2}](web-messaging.html#messageport)
    in `outside settings`{.variable}\'s
    [realm](webappapis.html#environment-settings-object's-realm){#dedicated-workers-and-the-worker-interface:environment-settings-object's-realm}.

7.  Set `outside port`{.variable}\'s [message event
    target](web-messaging.html#message-event-target){#dedicated-workers-and-the-worker-interface:message-event-target}
    to `worker`{.variable}.

8.  Set `worker`{.variable}\'s [outside
    port](#outside-port){#dedicated-workers-and-the-worker-interface:outside-port-2}
    to `outside port`{.variable}.

9.  Run this step [in
    parallel](infrastructure.html#in-parallel){#dedicated-workers-and-the-worker-interface:in-parallel}:

    1.  [Run a
        worker](#run-a-worker){#dedicated-workers-and-the-worker-interface:run-a-worker}
        given `worker`{.variable}, `worker URL`{.variable},
        `outside settings`{.variable}, `outside port`{.variable}, and
        `options`{.variable}.

10. Return `worker`{.variable}.

##### [10.2.6.4]{.secno} Shared workers and the [`SharedWorker`{#sharedworker-dev}](#sharedworker) interface[](#shared-workers-and-the-sharedworker-interface){.self-link}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[SharedWorker](https://developer.mozilla.org/en-US/docs/Web/API/SharedWorker "The SharedWorker interface represents a specific kind of worker that can be accessed from several browsing contexts, such as several windows, iframes or even workers. They implement an interface different than dedicated workers and have a different global scope, SharedWorkerGlobalScope.")

Support in all current engines.

::: support
[Firefox29+]{.firefox .yes}[Safari16+]{.safari .yes}[Chrome5+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android33+]{.firefox_android .yes}[Safari iOS16+]{.safari_ios
.yes}[Chrome AndroidNo]{.chrome_android .no}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet4.0--5.0]{.samsunginternet_android .no}[Opera
Android11--14]{.opera_android .no}
:::
::::
:::::

``` idl
[Exposed=Window]
interface SharedWorker : EventTarget {
  constructor((TrustedScriptURL or USVString) scriptURL, optional (DOMString or WorkerOptions) options = {});

  readonly attribute MessagePort port;
};
SharedWorker includes AbstractWorker;
```

`sharedWorker`{.variable}` = new `[`SharedWorker`](#dom-sharedworker){#dom-sharedworker-dev}`(``scriptURL`{.variable}` [, ``name`{.variable}` ])`

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[SharedWorker/SharedWorker](https://developer.mozilla.org/en-US/docs/Web/API/SharedWorker/SharedWorker "The SharedWorker() constructor creates a SharedWorker object that executes the script at the specified URL. This script must obey the same-origin policy.")

Support in all current engines.

::: support
[Firefox29+]{.firefox .yes}[Safari16+]{.safari .yes}[Chrome5+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android33+]{.firefox_android .yes}[Safari iOS16+]{.safari_ios
.yes}[Chrome AndroidNo]{.chrome_android .no}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet4.0--5.0]{.samsunginternet_android .no}[Opera
Android11--14]{.opera_android .no}
:::
::::
:::::

Returns a new
[`SharedWorker`{#shared-workers-and-the-sharedworker-interface:sharedworker-2}](#sharedworker)
object. `scriptURL`{.variable} will be fetched and executed in the
background, creating a new global environment for which
`sharedWorker`{.variable} represents the communication channel.
`name`{.variable} can be used to define the
[name](#concept-workerglobalscope-name){#shared-workers-and-the-sharedworker-interface:concept-workerglobalscope-name}
of that global environment.

`sharedWorker`{.variable}` = new `[`SharedWorker`](#dom-sharedworker){#shared-workers-and-the-sharedworker-interface:dom-sharedworker-2}`(``scriptURL`{.variable}` [, ``options`{.variable}` ])`

Returns a new
[`SharedWorker`{#shared-workers-and-the-sharedworker-interface:sharedworker-3}](#sharedworker)
object. `scriptURL`{.variable} will be fetched and executed in the
background, creating a new global environment for which
`sharedWorker`{.variable} represents the communication channel.
`options`{.variable} can be used to define the
[name](#concept-workerglobalscope-name){#shared-workers-and-the-sharedworker-interface:concept-workerglobalscope-name-2}
of that global environment via the `name` option. It can also ensure
this new global environment supports JavaScript modules (specify
`type: "module"`), and if that is specified, can also be used to specify
how `scriptURL`{.variable} is fetched through the `credentials` option.
Note that attempting to construct a shared worker with
`options`{.variable} whose `type` or `credentials` values mismatch an
existing shared worker will cause the returned `sharedWorker`{.variable}
to fire an error event and not connect to the existing shared worker.

`sharedWorker`{.variable}`.`[`port`](#dom-sharedworker-port){#dom-sharedworker-port-dev}

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[SharedWorker/port](https://developer.mozilla.org/en-US/docs/Web/API/SharedWorker/port "The port property of the SharedWorker interface returns a MessagePort object used to communicate and control the shared worker.")

Support in all current engines.

::: support
[Firefox29+]{.firefox .yes}[Safari16+]{.safari .yes}[Chrome5+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera10.6+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)?]{.edge .unknown}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android33+]{.firefox_android .yes}[Safari iOS16+]{.safari_ios
.yes}[Chrome AndroidNo]{.chrome_android .no}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet4.0--5.0]{.samsunginternet_android .no}[Opera
Android11--14]{.opera_android .no}
:::
::::
:::::

Returns `sharedWorker`{.variable}\'s
[`MessagePort`{#shared-workers-and-the-sharedworker-interface:messageport-2}](web-messaging.html#messageport)
object which can be used to communicate with the global environment.

The [`port`]{#dom-sharedworker-port .dfn dfn-for="SharedWorker"
dfn-type="attribute"} attribute must return the value it was assigned by
the object\'s constructor. It represents the
[`MessagePort`{#shared-workers-and-the-sharedworker-interface:messageport-3}](web-messaging.html#messageport)
for communicating with the shared worker.

A user agent has an associated [shared worker
manager]{#shared-worker-manager .dfn} which is the result of [starting a
new parallel
queue](infrastructure.html#starting-a-new-parallel-queue){#shared-workers-and-the-sharedworker-interface:starting-a-new-parallel-queue}.

Each user agent has a single [shared worker
manager](#shared-worker-manager){#shared-workers-and-the-sharedworker-interface:shared-worker-manager}
for simplicity. Implementations could use one per
[origin](browsers.html#concept-origin){#shared-workers-and-the-sharedworker-interface:concept-origin};
that would not be observably different and enables more concurrency.

When the
[`SharedWorker(``scriptURL`{.variable}`, ``options`{.variable}`)`]{#dom-sharedworker
.dfn dfn-for="SharedWorker" dfn-type="constructor"} constructor is
invoked:

1.  Let `compliantScriptURL`{.variable} be the result of invoking the
    [Get Trusted Type compliant
    string](https://w3c.github.io/trusted-types/dist/spec/#get-trusted-type-compliant-string-algorithm){#shared-workers-and-the-sharedworker-interface:tt-getcompliantstring
    x-internal="tt-getcompliantstring"} algorithm with
    [`TrustedScriptURL`{#shared-workers-and-the-sharedworker-interface:tt-trustedscripturl-2}](https://w3c.github.io/trusted-types/dist/spec/#trustedscripturl){x-internal="tt-trustedscripturl"},
    [this](https://webidl.spec.whatwg.org/#this){#shared-workers-and-the-sharedworker-interface:this
    x-internal="this"}\'s [relevant global
    object](webappapis.html#concept-relevant-global){#shared-workers-and-the-sharedworker-interface:concept-relevant-global},
    `scriptURL`{.variable}, \"`SharedWorker constructor`\", and
    \"`script`\".

2.  If `options`{.variable} is a
    [`DOMString`{#shared-workers-and-the-sharedworker-interface:idl-domstring}](https://webidl.spec.whatwg.org/#idl-DOMString){x-internal="idl-domstring"},
    set `options`{.variable} to a new
    [`WorkerOptions`{#shared-workers-and-the-sharedworker-interface:workeroptions-2}](#workeroptions)
    dictionary whose `name` member is set to the value of
    `options`{.variable} and whose other members are set to their
    default values.

3.  Let `outside settings`{.variable} be the [current settings
    object](webappapis.html#current-settings-object){#shared-workers-and-the-sharedworker-interface:current-settings-object}.

4.  Let `urlRecord`{.variable} be the result of [encoding-parsing a
    URL](urls-and-fetching.html#encoding-parsing-a-url){#shared-workers-and-the-sharedworker-interface:encoding-parsing-a-url}
    given `compliantScriptURL`{.variable}, relative to
    `outside settings`{.variable}.

    Any
    [same-origin](browsers.html#same-origin){#shared-workers-and-the-sharedworker-interface:same-origin}
    URL (including
    [`blob:`{#shared-workers-and-the-sharedworker-interface:blob-protocol}](https://w3c.github.io/FileAPI/#DefinitionOfScheme){x-internal="blob-protocol"}
    URLs) can be used.
    [`data:`{#shared-workers-and-the-sharedworker-interface:data-protocol}](https://www.rfc-editor.org/rfc/rfc2397#section-2){x-internal="data-protocol"}
    URLs can also be used, but they create a worker with an [opaque
    origin](browsers.html#concept-origin-opaque){#shared-workers-and-the-sharedworker-interface:concept-origin-opaque}.

5.  If `urlRecord`{.variable} is failure, then throw a
    [\"`SyntaxError`\"](https://webidl.spec.whatwg.org/#syntaxerror){#shared-workers-and-the-sharedworker-interface:syntaxerror
    x-internal="syntaxerror"}
    [`DOMException`{#shared-workers-and-the-sharedworker-interface:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

6.  Let `worker`{.variable} be a new
    [`SharedWorker`{#shared-workers-and-the-sharedworker-interface:sharedworker-4}](#sharedworker)
    object.

7.  Let `outside port`{.variable} be a
    [new](https://webidl.spec.whatwg.org/#new){#shared-workers-and-the-sharedworker-interface:new
    x-internal="new"}
    [`MessagePort`{#shared-workers-and-the-sharedworker-interface:messageport-4}](web-messaging.html#messageport)
    in `outside settings`{.variable}\'s
    [realm](webappapis.html#environment-settings-object's-realm){#shared-workers-and-the-sharedworker-interface:environment-settings-object's-realm}.

8.  Assign `outside port`{.variable} to the
    [`port`{#shared-workers-and-the-sharedworker-interface:dom-sharedworker-port-2}](#dom-sharedworker-port)
    attribute of `worker`{.variable}.

9.  Let `callerIsSecureContext`{.variable} be true if
    `outside settings`{.variable} is a [secure
    context](webappapis.html#secure-context){#shared-workers-and-the-sharedworker-interface:secure-context};
    otherwise, false.

10. Let `outside storage key`{.variable} be the result of running
    [obtain a storage key for non-storage
    purposes](https://storage.spec.whatwg.org/#obtain-a-storage-key-for-non-storage-purposes){#shared-workers-and-the-sharedworker-interface:obtain-a-storage-key-for-non-storage-purposes
    x-internal="obtain-a-storage-key-for-non-storage-purposes"} given
    `outside settings`{.variable}.

11. [Enqueue the following
    steps](infrastructure.html#enqueue-the-following-steps){#shared-workers-and-the-sharedworker-interface:enqueue-the-following-steps}
    to the [shared worker
    manager](#shared-worker-manager){#shared-workers-and-the-sharedworker-interface:shared-worker-manager-2}:

    1.  Let `worker global scope`{.variable} be null.

    2.  [For
        each](https://infra.spec.whatwg.org/#list-iterate){#shared-workers-and-the-sharedworker-interface:list-iterate
        x-internal="list-iterate"} `scope`{.variable} in the list of all
        [`SharedWorkerGlobalScope`{#shared-workers-and-the-sharedworker-interface:sharedworkerglobalscope}](#sharedworkerglobalscope)
        objects:

        1.  Let `worker storage key`{.variable} be the result of running
            [obtain a storage key for non-storage
            purposes](https://storage.spec.whatwg.org/#obtain-a-storage-key-for-non-storage-purposes){#shared-workers-and-the-sharedworker-interface:obtain-a-storage-key-for-non-storage-purposes-2
            x-internal="obtain-a-storage-key-for-non-storage-purposes"}
            given `scope`{.variable}\'s [relevant settings
            object](webappapis.html#relevant-settings-object){#shared-workers-and-the-sharedworker-interface:relevant-settings-object}.

        2.  If all of the following are true:

            - `worker storage key`{.variable}
              [equals](https://storage.spec.whatwg.org/#storage-key-equal){#shared-workers-and-the-sharedworker-interface:storage-key-equal
              x-internal="storage-key-equal"}
              `outside storage key`{.variable};
            - `scope`{.variable}\'s
              [closing](#dom-workerglobalscope-closing){#shared-workers-and-the-sharedworker-interface:dom-workerglobalscope-closing}
              flag is false;
            - `scope`{.variable}\'s [constructor
              url](#concept-sharedworkerglobalscope-constructor-url){#shared-workers-and-the-sharedworker-interface:concept-sharedworkerglobalscope-constructor-url}
              [equals](https://url.spec.whatwg.org/#concept-url-equals){#shared-workers-and-the-sharedworker-interface:concept-url-equals
              x-internal="concept-url-equals"} `urlRecord`{.variable};
              and
            - `scope`{.variable}\'s
              [name](#concept-workerglobalscope-name){#shared-workers-and-the-sharedworker-interface:concept-workerglobalscope-name-3}
              equals the value of `options`{.variable}\'s `name` member,

            then:

            1.  Set `worker global scope`{.variable} to
                `scope`{.variable}.

            2.  [Break](https://infra.spec.whatwg.org/#iteration-break){#shared-workers-and-the-sharedworker-interface:break
                x-internal="break"}.

        [`data:`{#shared-workers-and-the-sharedworker-interface:data-protocol-2}](https://www.rfc-editor.org/rfc/rfc2397#section-2){x-internal="data-protocol"}
        URLs create a worker with an [opaque
        origin](browsers.html#concept-origin-opaque){#shared-workers-and-the-sharedworker-interface:concept-origin-opaque-2}.
        Both the [constructor
        origin](#concept-sharedworkerglobalscope-constructor-origin){#shared-workers-and-the-sharedworker-interface:concept-sharedworkerglobalscope-constructor-origin}
        and [constructor
        url](#concept-sharedworkerglobalscope-constructor-url){#shared-workers-and-the-sharedworker-interface:concept-sharedworkerglobalscope-constructor-url-2}
        are compared so the same
        [`data:`{#shared-workers-and-the-sharedworker-interface:data-protocol-3}](https://www.rfc-editor.org/rfc/rfc2397#section-2){x-internal="data-protocol"}
        URL can be used within an
        [origin](browsers.html#concept-origin){#shared-workers-and-the-sharedworker-interface:concept-origin-2}
        to get to the same
        [`SharedWorkerGlobalScope`{#shared-workers-and-the-sharedworker-interface:sharedworkerglobalscope-2}](#sharedworkerglobalscope)
        object, but cannot be used to bypass the [same
        origin](browsers.html#same-origin){#shared-workers-and-the-sharedworker-interface:same-origin-2}
        restriction.

    3.  If `worker global scope`{.variable} is not null, but the user
        agent has been configured to disallow communication between the
        worker represented by the `worker global scope`{.variable} and
        the
        [scripts](webappapis.html#concept-script){#shared-workers-and-the-sharedworker-interface:concept-script}
        whose [settings
        object](webappapis.html#settings-object){#shared-workers-and-the-sharedworker-interface:settings-object}
        is `outside settings`{.variable}, then set
        `worker global scope`{.variable} to null.

        For example, a user agent could have a development mode that
        isolates a particular [top-level
        traversable](document-sequences.html#top-level-traversable){#shared-workers-and-the-sharedworker-interface:top-level-traversable}
        from all other pages, and scripts in that development mode could
        be blocked from connecting to shared workers running in the
        normal browser mode.

    4.  If `worker global scope`{.variable} is not null, then check if
        `worker global scope`{.variable}\'s
        [type](#concept-workerglobalscope-type){#shared-workers-and-the-sharedworker-interface:concept-workerglobalscope-type}
        and
        [credentials](#concept-sharedworkerglobalscope-credentials){#shared-workers-and-the-sharedworker-interface:concept-sharedworkerglobalscope-credentials}
        match the `options`{.variable} values. If not, [queue a
        task](webappapis.html#queue-a-task){#shared-workers-and-the-sharedworker-interface:queue-a-task}
        to [fire an
        event](https://dom.spec.whatwg.org/#concept-event-fire){#shared-workers-and-the-sharedworker-interface:concept-event-fire
        x-internal="concept-event-fire"} named
        [`error`{#shared-workers-and-the-sharedworker-interface:event-error}](indices.html#event-error)
        and abort these steps.

    5.  If `worker global scope`{.variable} is not null, then run these
        subsubsteps:

        1.  Let `settings object`{.variable} be the [relevant settings
            object](webappapis.html#relevant-settings-object){#shared-workers-and-the-sharedworker-interface:relevant-settings-object-2}
            for `worker global scope`{.variable}.

        2.  Let `workerIsSecureContext`{.variable} be true if
            `settings object`{.variable} is a [secure
            context](webappapis.html#secure-context){#shared-workers-and-the-sharedworker-interface:secure-context-2};
            otherwise, false.

        3.  If `workerIsSecureContext`{.variable} is not
            `callerIsSecureContext`{.variable}, then [queue a
            task](webappapis.html#queue-a-task){#shared-workers-and-the-sharedworker-interface:queue-a-task-2}
            to [fire an
            event](https://dom.spec.whatwg.org/#concept-event-fire){#shared-workers-and-the-sharedworker-interface:concept-event-fire-2
            x-internal="concept-event-fire"} named
            [`error`{#shared-workers-and-the-sharedworker-interface:event-error-2}](indices.html#event-error)
            at `worker`{.variable} and abort these steps.
            [\[SECURE-CONTEXTS\]](references.html#refsSECURE-CONTEXTS)

        4.  Associate `worker`{.variable} with
            `worker global scope`{.variable}.

        5.  Let `inside port`{.variable} be a
            [new](https://webidl.spec.whatwg.org/#new){#shared-workers-and-the-sharedworker-interface:new-2
            x-internal="new"}
            [`MessagePort`{#shared-workers-and-the-sharedworker-interface:messageport-5}](web-messaging.html#messageport)
            in `settings object`{.variable}\'s
            [realm](webappapis.html#environment-settings-object's-realm){#shared-workers-and-the-sharedworker-interface:environment-settings-object's-realm-2}.

        6.  [Entangle](web-messaging.html#entangle){#shared-workers-and-the-sharedworker-interface:entangle}
            `outside port`{.variable} and `inside port`{.variable}.

        7.  [Queue a
            task](webappapis.html#queue-a-task){#shared-workers-and-the-sharedworker-interface:queue-a-task-3},
            using the [DOM manipulation task
            source](webappapis.html#dom-manipulation-task-source){#shared-workers-and-the-sharedworker-interface:dom-manipulation-task-source},
            to [fire an
            event](https://dom.spec.whatwg.org/#concept-event-fire){#shared-workers-and-the-sharedworker-interface:concept-event-fire-3
            x-internal="concept-event-fire"} named
            [`connect`{#shared-workers-and-the-sharedworker-interface:event-workerglobalscope-connect}](indices.html#event-workerglobalscope-connect)
            at `worker global scope`{.variable}, using
            [`MessageEvent`{#shared-workers-and-the-sharedworker-interface:messageevent}](comms.html#messageevent),
            with the
            [`data`{#shared-workers-and-the-sharedworker-interface:dom-messageevent-data}](comms.html#dom-messageevent-data)
            attribute initialized to the empty string, the
            [`ports`{#shared-workers-and-the-sharedworker-interface:dom-messageevent-ports}](comms.html#dom-messageevent-ports)
            attribute initialized to a new [frozen
            array](https://webidl.spec.whatwg.org/#dfn-frozen-array-type){#shared-workers-and-the-sharedworker-interface:frozen-array
            x-internal="frozen-array"} containing only
            `inside port`{.variable}, and the
            [`source`{#shared-workers-and-the-sharedworker-interface:dom-messageevent-source}](comms.html#dom-messageevent-source)
            attribute initialized to `inside port`{.variable}.

        8.  [Append](https://infra.spec.whatwg.org/#set-append){#shared-workers-and-the-sharedworker-interface:set-append
            x-internal="set-append"} the [relevant owner to
            add](#relevant-owner-to-add){#shared-workers-and-the-sharedworker-interface:relevant-owner-to-add}
            given `outside settings`{.variable} to
            `worker global scope`{.variable}\'s [owner
            set](#concept-WorkerGlobalScope-owner-set){#shared-workers-and-the-sharedworker-interface:concept-WorkerGlobalScope-owner-set}.

    6.  Otherwise, [in
        parallel](infrastructure.html#in-parallel){#shared-workers-and-the-sharedworker-interface:in-parallel},
        [run a
        worker](#run-a-worker){#shared-workers-and-the-sharedworker-interface:run-a-worker}
        given `worker`{.variable}, `urlRecord`{.variable},
        `outside settings`{.variable}, `outside port`{.variable}, and
        `options`{.variable}.

12. Return `worker`{.variable}.

#### [10.2.7]{.secno} Concurrent hardware capabilities[](#navigator.hardwareconcurrency){.self-link} {#navigator.hardwareconcurrency}

``` idl
interface mixin NavigatorConcurrentHardware {
  readonly attribute unsigned long long hardwareConcurrency;
};
```

`self`{.variable}`.`[`navigator`](system-state.html#dom-navigator){#navigator.hardwareconcurrency:dom-navigator}`.`[`hardwareConcurrency`](#dom-navigator-hardwareconcurrency){#dom-navigator-hardwareconcurrency-dev}

::::::: {.mdn-anno .wrapped .before}
MDN

:::: feature
[Navigator/hardwareConcurrency](https://developer.mozilla.org/en-US/docs/Web/API/Navigator/hardwareConcurrency "The navigator.hardwareConcurrency read-only property returns the number of logical processors available to run threads on the user's computer.")

::: support
[Firefox48+]{.firefox .yes}[Safari10.1--11]{.safari
.no}[Chrome37+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)15+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::

:::: feature
[Navigator/hardwareConcurrency](https://developer.mozilla.org/en-US/docs/Web/API/Navigator/hardwareConcurrency "The navigator.hardwareConcurrency read-only property returns the number of logical processors available to run threads on the user's computer.")

::: support
[Firefox48+]{.firefox .yes}[Safari10.1--11]{.safari
.no}[Chrome37+]{.chrome .yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)15+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::::

Returns the number of logical processors potentially available to the
user agent.

[![(This is a tracking
vector.)](https://resources.whatwg.org/tracking-vector.svg){width="46"
crossorigin=""
height="64"}](https://infra.spec.whatwg.org/#tracking-vector "There is a tracking vector here."){#navigator.hardwareconcurrency:tracking-vector
.tracking-vector x-internal="tracking-vector"} The
[`navigator.hardwareConcurrency`]{#dom-navigator-hardwareconcurrency
.dfn dfn-for="NavigatorConcurrentHardware" dfn-type="attribute"}
attribute\'s getter must return a number between 1 and the number of
logical processors potentially available to the user agent. If this
cannot be determined, the getter must return 1.

User agents should err toward exposing the number of logical processors
available, using lower values only in cases where there are user-agent
specific limits in place (such as a limitation on the number of
[workers](#worker){#navigator.hardwareconcurrency:worker} that can be
created) or when the user agent desires to limit fingerprinting
possibilities.

### [10.3]{.secno} APIs available to workers[](#apis-available-to-workers){.self-link}

#### [10.3.1]{.secno} Importing scripts and libraries[](#importing-scripts-and-libraries){.self-link}

The
[`importScripts(...``urls`{.variable}`)`]{#dom-workerglobalscope-importscripts
.dfn dfn-for="WorkerGlobalScope" dfn-type="method"} method steps are:

1.  Let `urlStrings`{.variable} be « ».

2.  [For
    each](https://infra.spec.whatwg.org/#list-iterate){#importing-scripts-and-libraries:list-iterate
    x-internal="list-iterate"} `url`{.variable} of `urls`{.variable}:

    1.  [Append](https://infra.spec.whatwg.org/#list-append){#importing-scripts-and-libraries:list-append
        x-internal="list-append"} the result of invoking the [Get
        Trusted Type compliant
        string](https://w3c.github.io/trusted-types/dist/spec/#get-trusted-type-compliant-string-algorithm){#importing-scripts-and-libraries:tt-getcompliantstring
        x-internal="tt-getcompliantstring"} algorithm with
        [`TrustedScriptURL`{#importing-scripts-and-libraries:tt-trustedscripturl}](https://w3c.github.io/trusted-types/dist/spec/#trustedscripturl){x-internal="tt-trustedscripturl"},
        [this](https://webidl.spec.whatwg.org/#this){#importing-scripts-and-libraries:this
        x-internal="this"}\'s [relevant global
        object](webappapis.html#concept-relevant-global){#importing-scripts-and-libraries:concept-relevant-global},
        `url`{.variable}, \"`WorkerGlobalScope importScripts`\", and
        \"`script`\" to `urlStrings`{.variable}.

3.  [Import scripts into worker global
    scope](#import-scripts-into-worker-global-scope){#importing-scripts-and-libraries:import-scripts-into-worker-global-scope}
    given
    [this](https://webidl.spec.whatwg.org/#this){#importing-scripts-and-libraries:this-2
    x-internal="this"} and `urlStrings`{.variable}.

To [import scripts into worker global
scope]{#import-scripts-into-worker-global-scope .dfn export=""}, given a
[`WorkerGlobalScope`{#importing-scripts-and-libraries:workerglobalscope}](#workerglobalscope)
object `worker global scope`{.variable}, a
[list](https://infra.spec.whatwg.org/#list){#importing-scripts-and-libraries:list
x-internal="list"} of [scalar value
strings](https://infra.spec.whatwg.org/#scalar-value-string){#importing-scripts-and-libraries:scalar-value-string
x-internal="scalar-value-string"} `urls`{.variable}, and an optional
[perform the fetch
hook](webappapis.html#fetching-scripts-perform-fetch){#importing-scripts-and-libraries:fetching-scripts-perform-fetch}
`performFetch`{.variable}:

1.  If `worker global scope`{.variable}\'s
    [type](#concept-workerglobalscope-type){#importing-scripts-and-libraries:concept-workerglobalscope-type}
    is \"`module`\", throw a
    [`TypeError`{#importing-scripts-and-libraries:typeerror}](https://tc39.es/ecma262/#sec-native-error-types-used-in-this-standard-typeerror){x-internal="typeerror"}
    exception.

2.  Let `settings object`{.variable} be the [current settings
    object](webappapis.html#current-settings-object){#importing-scripts-and-libraries:current-settings-object}.

3.  If `urls`{.variable} is empty, return.

4.  Let `urlRecords`{.variable} be « ».

5.  For each `url`{.variable} of `urls`{.variable}:

    1.  Let `urlRecord`{.variable} be the result of [encoding-parsing a
        URL](urls-and-fetching.html#encoding-parsing-a-url){#importing-scripts-and-libraries:encoding-parsing-a-url}
        given `url`{.variable}, relative to
        `settings object`{.variable}.

    2.  If `urlRecord`{.variable} is failure, then throw a
        [\"`SyntaxError`\"](https://webidl.spec.whatwg.org/#syntaxerror){#importing-scripts-and-libraries:syntaxerror
        x-internal="syntaxerror"}
        [`DOMException`{#importing-scripts-and-libraries:domexception}](https://webidl.spec.whatwg.org/#dfn-DOMException){x-internal="domexception"}.

    3.  [Append](https://infra.spec.whatwg.org/#list-append){#importing-scripts-and-libraries:list-append-2
        x-internal="list-append"} `urlRecord`{.variable} to
        `urlRecords`{.variable}.

6.  For each `urlRecord`{.variable} of `urlRecords`{.variable}:

    1.  [Fetch a classic worker-imported
        script](webappapis.html#fetch-a-classic-worker-imported-script){#importing-scripts-and-libraries:fetch-a-classic-worker-imported-script}
        given `urlRecord`{.variable} and `settings object`{.variable},
        passing along `performFetch`{.variable} if provided. If this
        succeeds, let `script`{.variable} be the result. Otherwise,
        rethrow the exception.

    2.  [Run the classic
        script](webappapis.html#run-a-classic-script){#importing-scripts-and-libraries:run-a-classic-script}
        `script`{.variable}, with *rethrow errors* set to true.

        `script`{.variable} will run until it either returns, fails to
        parse, fails to catch an exception, or gets [prematurely
        aborted](webappapis.html#abort-a-running-script){#importing-scripts-and-libraries:abort-a-running-script}
        by the [terminate a
        worker](#terminate-a-worker){#importing-scripts-and-libraries:terminate-a-worker}
        algorithm defined above.

        If an exception was thrown or if the script was [prematurely
        aborted](webappapis.html#abort-a-running-script){#importing-scripts-and-libraries:abort-a-running-script-2},
        then abort all these steps, letting the exception or aborting
        continue to be processed by the calling
        [script](webappapis.html#concept-script){#importing-scripts-and-libraries:concept-script}.

Service Workers is an example of a specification that runs this
algorithm with its own [perform the fetch
hook](webappapis.html#fetching-scripts-perform-fetch){#importing-scripts-and-libraries:fetching-scripts-perform-fetch-2}.
[\[SW\]](references.html#refsSW)

#### [10.3.2]{.secno} The [`WorkerNavigator`{#workernavigator-dev}](#workernavigator) interface[](#the-workernavigator-object){.self-link} {#the-workernavigator-object}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[WorkerNavigator](https://developer.mozilla.org/en-US/docs/Web/API/WorkerNavigator "The WorkerNavigator interface represents a subset of the Navigator interface allowed to be accessed from a Worker. Such an object is initialized for each worker and is available via the self.navigator property.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The [`navigator`]{#dom-worker-navigator .dfn dfn-for="WorkerGlobalScope"
dfn-type="attribute"} attribute of the
[`WorkerGlobalScope`{#the-workernavigator-object:workerglobalscope}](#workerglobalscope)
interface must return an instance of the
[`WorkerNavigator`{#the-workernavigator-object:workernavigator}](#workernavigator)
interface, which represents the identity and state of the user agent
(the client):

``` idl
[Exposed=Worker]
interface WorkerNavigator {};
WorkerNavigator includes NavigatorID;
WorkerNavigator includes NavigatorLanguage;
WorkerNavigator includes NavigatorOnLine;
WorkerNavigator includes NavigatorConcurrentHardware;
```

#### [10.3.3]{.secno} The [`WorkerLocation`{#workerlocation-dev}](#workerlocation) interface[](#worker-locations){.self-link} {#worker-locations}

::::: {.mdn-anno .wrapped}
**✔**MDN

:::: feature
[WorkerLocation](https://developer.mozilla.org/en-US/docs/Web/API/WorkerLocation "The WorkerLocation interface defines the absolute location of the script executed by the Worker. Such an object is initialized for each worker and is available via the WorkerGlobalScope.location property obtained by calling self.location.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[WorkerLocation/toString](https://developer.mozilla.org/en-US/docs/Web/API/WorkerLocation/toString "The toString() stringifier method of a WorkerLocation object returns a string containing the serialized URL for the worker's location. It is a synonym for WorkerLocation.href.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera15+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android37+]{.webview_android .yes}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android14+]{.opera_android .yes}
:::
::::
:::::

``` idl
[Exposed=Worker]
interface WorkerLocation {
  stringifier readonly attribute USVString href;
  readonly attribute USVString origin;
  readonly attribute USVString protocol;
  readonly attribute USVString host;
  readonly attribute USVString hostname;
  readonly attribute USVString port;
  readonly attribute USVString pathname;
  readonly attribute USVString search;
  readonly attribute USVString hash;
};
```

A [`WorkerLocation`{#worker-locations:workerlocation}](#workerlocation)
object has an associated [`WorkerGlobalScope`
object]{#concept-workerlocation-workerglobalscope .dfn} (a
[`WorkerGlobalScope`{#worker-locations:workerglobalscope}](#workerglobalscope)
object).

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[WorkerLocation/href](https://developer.mozilla.org/en-US/docs/Web/API/WorkerLocation/href "The href property of a WorkerLocation object returns a string containing the serialized URL for the worker's location.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The [`href`]{#dom-workerlocation-href .dfn dfn-for="WorkerLocation"
dfn-type="attribute"} getter steps are to return
[this](https://webidl.spec.whatwg.org/#this){#worker-locations:this
x-internal="this"}\'s [`WorkerGlobalScope`
object](#concept-workerlocation-workerglobalscope){#worker-locations:concept-workerlocation-workerglobalscope}\'s
[url](#concept-workerglobalscope-url){#worker-locations:concept-workerglobalscope-url},
[serialized](https://url.spec.whatwg.org/#concept-url-serializer){#worker-locations:concept-url-serializer
x-internal="concept-url-serializer"}.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[WorkerLocation/origin](https://developer.mozilla.org/en-US/docs/Web/API/WorkerLocation/origin "The origin property of a WorkerLocation object returns the worker's origin.")

Support in all current engines.

::: support
[Firefox29+]{.firefox .yes}[Safari10+]{.safari .yes}[Chrome38+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera?]{.opera .unknown}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)14+]{.edge .yes}[Internet ExplorerNo]{.ie .no}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS?]{.safari_ios
.unknown}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android?]{.opera_android .unknown}
:::
::::
:::::

The [`origin`]{#dom-workerlocation-origin .dfn dfn-for="WorkerLocation"
dfn-type="attribute"} getter steps are to return the
[serialization](browsers.html#ascii-serialisation-of-an-origin){#worker-locations:ascii-serialisation-of-an-origin}
of [this](https://webidl.spec.whatwg.org/#this){#worker-locations:this-2
x-internal="this"}\'s [`WorkerGlobalScope`
object](#concept-workerlocation-workerglobalscope){#worker-locations:concept-workerlocation-workerglobalscope-2}\'s
[url](#concept-workerglobalscope-url){#worker-locations:concept-workerglobalscope-url-2}\'s
[origin](https://url.spec.whatwg.org/#concept-url-origin){#worker-locations:concept-url-origin
x-internal="concept-url-origin"}.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[WorkerLocation/protocol](https://developer.mozilla.org/en-US/docs/Web/API/WorkerLocation/protocol "The protocol property of a WorkerLocation object returns the protocol part of the worker's location.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The [`protocol`]{#dom-workerlocation-protocol .dfn
dfn-for="WorkerLocation" dfn-type="attribute"} getter steps are to
return
[this](https://webidl.spec.whatwg.org/#this){#worker-locations:this-3
x-internal="this"}\'s [`WorkerGlobalScope`
object](#concept-workerlocation-workerglobalscope){#worker-locations:concept-workerlocation-workerglobalscope-3}\'s
[url](#concept-workerglobalscope-url){#worker-locations:concept-workerglobalscope-url-3}\'s
[scheme](https://url.spec.whatwg.org/#concept-url-scheme){#worker-locations:concept-url-scheme
x-internal="concept-url-scheme"}, followed by \"`:`\".

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[WorkerLocation/host](https://developer.mozilla.org/en-US/docs/Web/API/WorkerLocation/host "The host property of a WorkerLocation object returns the host part of the worker's location.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The [`host`]{#dom-workerlocation-host .dfn dfn-for="WorkerLocation"
dfn-type="attribute"} getter steps are:

1.  Let `url`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#worker-locations:this-4
    x-internal="this"}\'s [`WorkerGlobalScope`
    object](#concept-workerlocation-workerglobalscope){#worker-locations:concept-workerlocation-workerglobalscope-4}\'s
    [url](#concept-workerglobalscope-url){#worker-locations:concept-workerglobalscope-url-4}.

2.  If `url`{.variable}\'s
    [host](https://url.spec.whatwg.org/#concept-url-host){#worker-locations:concept-url-host
    x-internal="concept-url-host"} is null, return the empty string.

3.  If `url`{.variable}\'s
    [port](https://url.spec.whatwg.org/#concept-url-port){#worker-locations:concept-url-port
    x-internal="concept-url-port"} is null, return `url`{.variable}\'s
    [host](https://url.spec.whatwg.org/#concept-url-host){#worker-locations:concept-url-host-2
    x-internal="concept-url-host"},
    [serialized](https://url.spec.whatwg.org/#concept-host-serializer){#worker-locations:host-serializer
    x-internal="host-serializer"}.

4.  Return `url`{.variable}\'s
    [host](https://url.spec.whatwg.org/#concept-url-host){#worker-locations:concept-url-host-3
    x-internal="concept-url-host"},
    [serialized](https://url.spec.whatwg.org/#concept-host-serializer){#worker-locations:host-serializer-2
    x-internal="host-serializer"}, followed by \"`:`\" and
    `url`{.variable}\'s
    [port](https://url.spec.whatwg.org/#concept-url-port){#worker-locations:concept-url-port-2
    x-internal="concept-url-port"},
    [serialized](https://url.spec.whatwg.org/#serialize-an-integer){#worker-locations:serialize-an-integer
    x-internal="serialize-an-integer"}.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[WorkerLocation/hostname](https://developer.mozilla.org/en-US/docs/Web/API/WorkerLocation/hostname "The hostname property of a WorkerLocation object returns the hostname part of the worker's location.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The [`hostname`]{#dom-workerlocation-hostname .dfn
dfn-for="WorkerLocation" dfn-type="attribute"} getter steps are:

1.  Let `host`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#worker-locations:this-5
    x-internal="this"}\'s [`WorkerGlobalScope`
    object](#concept-workerlocation-workerglobalscope){#worker-locations:concept-workerlocation-workerglobalscope-5}\'s
    [url](#concept-workerglobalscope-url){#worker-locations:concept-workerglobalscope-url-5}\'s
    [host](https://url.spec.whatwg.org/#concept-url-host){#worker-locations:concept-url-host-4
    x-internal="concept-url-host"}.

2.  If `host`{.variable} is null, return the empty string.

3.  Return `host`{.variable},
    [serialized](https://url.spec.whatwg.org/#concept-host-serializer){#worker-locations:host-serializer-3
    x-internal="host-serializer"}.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[WorkerLocation/port](https://developer.mozilla.org/en-US/docs/Web/API/WorkerLocation/port "The port property of a WorkerLocation object returns the port part of the worker's location.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The [`port`]{#dom-workerlocation-port .dfn dfn-for="WorkerLocation"
dfn-type="attribute"} getter steps are:

1.  Let `port`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#worker-locations:this-6
    x-internal="this"}\'s [`WorkerGlobalScope`
    object](#concept-workerlocation-workerglobalscope){#worker-locations:concept-workerlocation-workerglobalscope-6}\'s
    [url](#concept-workerglobalscope-url){#worker-locations:concept-workerglobalscope-url-6}\'s
    [port](https://url.spec.whatwg.org/#concept-url-port){#worker-locations:concept-url-port-3
    x-internal="concept-url-port"}.

2.  If `port`{.variable} is null, return the empty string.

3.  Return `port`{.variable},
    [serialized](https://url.spec.whatwg.org/#serialize-an-integer){#worker-locations:serialize-an-integer-2
    x-internal="serialize-an-integer"}.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[WorkerLocation/pathname](https://developer.mozilla.org/en-US/docs/Web/API/WorkerLocation/pathname "The pathname property of a WorkerLocation object returns the pathname part of the worker's location.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The [`pathname`]{#dom-workerlocation-pathname .dfn
dfn-for="WorkerLocation" dfn-type="attribute"} getter steps are to
return the result of [URL path
serializing](https://url.spec.whatwg.org/#url-path-serializer){#worker-locations:url-path-serializer
x-internal="url-path-serializer"}
[this](https://webidl.spec.whatwg.org/#this){#worker-locations:this-7
x-internal="this"}\'s [`WorkerGlobalScope`
object](#concept-workerlocation-workerglobalscope){#worker-locations:concept-workerlocation-workerglobalscope-7}\'s
[url](#concept-workerglobalscope-url){#worker-locations:concept-workerglobalscope-url-7}.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[WorkerLocation/search](https://developer.mozilla.org/en-US/docs/Web/API/WorkerLocation/search "The search property of a WorkerLocation object returns the search part of the worker's location.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The [`search`]{#dom-workerlocation-search .dfn dfn-for="WorkerLocation"
dfn-type="attribute"} getter steps are:

1.  Let `query`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#worker-locations:this-8
    x-internal="this"}\'s [`WorkerGlobalScope`
    object](#concept-workerlocation-workerglobalscope){#worker-locations:concept-workerlocation-workerglobalscope-8}\'s
    [url](#concept-workerglobalscope-url){#worker-locations:concept-workerglobalscope-url-8}\'s
    [query](https://url.spec.whatwg.org/#concept-url-query){#worker-locations:concept-url-query
    x-internal="concept-url-query"}.

2.  If `query`{.variable} is either null or the empty string, return the
    empty string.

3.  Return \"`?`\", followed by `query`{.variable}.

::::: {.mdn-anno .wrapped .before}
**✔**MDN

:::: feature
[WorkerLocation/hash](https://developer.mozilla.org/en-US/docs/Web/API/WorkerLocation/hash "The hash property of a WorkerLocation object returns the hash part of the worker's location.")

Support in all current engines.

::: support
[Firefox3.5+]{.firefox .yes}[Safari4+]{.safari .yes}[Chrome4+]{.chrome
.yes}

------------------------------------------------------------------------

[Opera12.1+]{.opera .yes}[Edge79+]{.edge_blink .yes}

------------------------------------------------------------------------

[Edge (Legacy)12+]{.edge .yes}[Internet Explorer10+]{.ie .yes}

------------------------------------------------------------------------

[Firefox Android?]{.firefox_android .unknown}[Safari iOS5+]{.safari_ios
.yes}[Chrome Android?]{.chrome_android .unknown}[WebView
Android?]{.webview_android .unknown}[Samsung
Internet?]{.samsunginternet_android .unknown}[Opera
Android12.1+]{.opera_android .yes}
:::
::::
:::::

The [`hash`]{#dom-workerlocation-hash .dfn dfn-for="WorkerLocation"
dfn-type="attribute"} getter steps are:

1.  Let `fragment`{.variable} be
    [this](https://webidl.spec.whatwg.org/#this){#worker-locations:this-9
    x-internal="this"}\'s [`WorkerGlobalScope`
    object](#concept-workerlocation-workerglobalscope){#worker-locations:concept-workerlocation-workerglobalscope-9}\'s
    [url](#concept-workerglobalscope-url){#worker-locations:concept-workerglobalscope-url-9}\'s
    [fragment](https://url.spec.whatwg.org/#concept-url-fragment){#worker-locations:concept-url-fragment
    x-internal="concept-url-fragment"}.

2.  If `fragment`{.variable} is either null or the empty string, return
    the empty string.

3.  Return \"`#`\", followed by `fragment`{.variable}.

[← 9.3 Cross-document messaging](web-messaging.html) --- [Table of
Contents](./) --- [11 Worklets →](worklets.html)
