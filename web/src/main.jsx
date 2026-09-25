import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App'
import './styles.css'

if (typeof window !== 'undefined') {
  const isPackaged = window.location.protocol === 'file:' || window.location.protocol === 'wails:' || !window.location.hostname
  window.API_BASE = isPackaged ? 'http://127.0.0.1:8080' : ''

  window.addEventListener('error', (ev) => {
    try {
      console.error('Window error:', ev.message, ev.error)
    } catch (e) {}
  })

  window.addEventListener('unhandledrejection', (ev) => {
    try {
      console.error('Unhandled promise rejection:', ev.reason)
    } catch (e) {}
  })

  const _fetch = window.fetch.bind(window)
  window.fetch = (input, init) => {
    let originalInput = input
    try {
      if (typeof input === 'string') {
        if (input.startsWith('/')) input = (window.API_BASE || '') + input
      } else if (input && input.url && typeof input.url === 'string' && input.url.startsWith('/')) {
        input = new Request((window.API_BASE || '') + input.url, input)
      }
    } catch (e) {
      console.error('Error rewriting fetch URL:', e)
    }

    try {
      if (isPackaged) {
        const method = (init && init.method) || (input && input.method) || 'GET'
        const url = typeof input === 'string' ? input : input.url
        console.log('[fetch] ', method, url, init && init.body ? init.body : '')
      }
    } catch (e) {}

    return _fetch(input, init)
  }

  if (isPackaged) {
    try {
      _fetch((window.API_BASE || '') + '/health')
        .then((r) => console.log('[health] status', r.status))
        .catch((err) => console.error('[health] fetch error', err))
    } catch (e) {
      console.error('Health check failed', e)
    }
  }
}

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
)
