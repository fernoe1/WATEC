import React, { useMemo, useState } from 'react';
import { createRoot } from 'react-dom/client';
import { Plus, Search, Save, Trash2, RefreshCw, Clock, DoorOpen, Server } from 'lucide-react';
import './styles.css';

const DEFAULT_API_BASE = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api/v1';

const emptySlot = () => ({ roomNumber: '', from: '', to: '' });

function normalizeFreeSlot(slot) {
  return {
    roomNumber: Number(slot.roomNumber),
    from: Number(slot.from),
    to: Number(slot.to),
  };
}

async function requestJson(url, options = {}) {
  const response = await fetch(url, {
    headers: {
      'Content-Type': 'application/json',
      ...(options.headers || {}),
    },
    ...options,
  });

  let data = null;
  const text = await response.text();
  if (text) {
    try {
      data = JSON.parse(text);
    } catch {
      data = { raw: text };
    }
  }

  if (!response.ok) {
    const message = data?.error || data?.message || `Request failed with status ${response.status}`;
    throw new Error(message);
  }

  return data;
}

function App() {
  const [apiBase, setApiBase] = useState(DEFAULT_API_BASE);
  const [name, setName] = useState('');
  const [free, setFree] = useState([emptySlot()]);
  const [searchName, setSearchName] = useState('');
  const [result, setResult] = useState(null);
  const [message, setMessage] = useState(null);
  const [loading, setLoading] = useState(false);

  const teacherUrl = useMemo(() => `${apiBase.replace(/\/$/, '')}/teachers`, [apiBase]);

  const showMessage = (type, text) => setMessage({ type, text });

  const payload = () => ({
    name: name.trim(),
    free: free
      .filter((slot) => slot.roomNumber !== '' && slot.from !== '' && slot.to !== '')
      .map(normalizeFreeSlot),
  });

  const validateWrite = () => {
    const body = payload();
    if (!body.name) throw new Error('Teacher name is required');
    for (const slot of body.free) {
      if (slot.from < 0 || slot.from > 23 || slot.to < 1 || slot.to > 24 || slot.from >= slot.to) {
        throw new Error('Each free slot must have valid hours: from 0-23, to 1-24, and from < to');
      }
      if (slot.roomNumber <= 0) throw new Error('Room number must be positive');
    }
    return body;
  };

  const run = async (action) => {
    setLoading(true);
    setMessage(null);
    try {
      await action();
    } catch (error) {
      showMessage('error', error.message);
    } finally {
      setLoading(false);
    }
  };

  const createTeacher = () => run(async () => {
    const body = validateWrite();
    const data = await requestJson(teacherUrl, {
      method: 'POST',
      body: JSON.stringify(body),
    });
    showMessage('success', data?.message || 'Teacher created');
  });

  const updateTeacher = () => run(async () => {
    const body = validateWrite();
    const data = await requestJson(`${teacherUrl}/${encodeURIComponent(body.name)}`, {
      method: 'PUT',
      body: JSON.stringify(body),
    });
    showMessage('success', data?.message || 'Teacher updated');
  });

  const readTeacher = () => run(async () => {
    const target = searchName.trim() || name.trim();
    if (!target) throw new Error('Enter teacher name to search');
    const data = await requestJson(`${teacherUrl}/${encodeURIComponent(target)}`, { method: 'GET' });
    setResult({ name: target, ...data });
    showMessage('success', 'Teacher availability loaded');
  });

  const deleteTeacher = () => run(async () => {
    const target = searchName.trim() || name.trim();
    if (!target) throw new Error('Enter teacher name to delete');
    const data = await requestJson(`${teacherUrl}/${encodeURIComponent(target)}`, { method: 'DELETE' });
    setResult(null);
    showMessage('success', data?.message || 'Teacher deleted');
  });

  const addSlot = () => setFree((items) => [...items, emptySlot()]);
  const removeSlot = (index) => setFree((items) => items.filter((_, i) => i !== index));
  const updateSlot = (index, field, value) => {
    setFree((items) => items.map((slot, i) => (i === index ? { ...slot, [field]: value } : slot)));
  };

  const loadDemo = () => {
    setName('Aidos');
    setSearchName('Aidos');
    setFree([
      { roomNumber: 204, from: 9, to: 12 },
      { roomNumber: 305, from: 14, to: 17 },
    ]);
    setResult(null);
    setMessage(null);
  };

  return (
    <main className="page">
      <section className="hero">
        <div>
          <p className="eyebrow">WATEC Frontend</p>
          <h1>Teacher Availability Dashboard</h1>
          <p className="subtitle">Create, update, search and delete teacher schedules through the Gateway HTTP API.</p>
        </div>
        <div className="hero-card">
          <Server size={22} />
          <span>Gateway API</span>
          <input value={apiBase} onChange={(event) => setApiBase(event.target.value)} />
        </div>
      </section>

      {message && <div className={`alert ${message.type}`}>{message.text}</div>}

      <section className="grid">
        <div className="card wide">
          <div className="card-header">
            <div>
              <h2>Teacher form</h2>
              <p>Add name and hourly free slots.</p>
            </div>
            <button className="secondary" onClick={loadDemo} type="button">Demo data</button>
          </div>

          <label className="field">
            <span>Teacher name</span>
            <input placeholder="Example: Aidos" value={name} onChange={(event) => setName(event.target.value)} />
          </label>

          <div className="slots-title">
            <h3>Free slots</h3>
            <button className="ghost" onClick={addSlot} type="button"><Plus size={16} /> Add slot</button>
          </div>

          <div className="slots">
            {free.map((slot, index) => (
              <div className="slot" key={index}>
                <label>
                  <DoorOpen size={16} />
                  <input type="number" min="1" placeholder="Room" value={slot.roomNumber} onChange={(event) => updateSlot(index, 'roomNumber', event.target.value)} />
                </label>
                <label>
                  <Clock size={16} />
                  <input type="number" min="0" max="23" placeholder="From" value={slot.from} onChange={(event) => updateSlot(index, 'from', event.target.value)} />
                </label>
                <label>
                  <Clock size={16} />
                  <input type="number" min="1" max="24" placeholder="To" value={slot.to} onChange={(event) => updateSlot(index, 'to', event.target.value)} />
                </label>
                <button className="icon danger" onClick={() => removeSlot(index)} disabled={free.length === 1} type="button">
                  <Trash2 size={16} />
                </button>
              </div>
            ))}
          </div>

          <div className="actions">
            <button onClick={createTeacher} disabled={loading} type="button"><Plus size={18} /> Create</button>
            <button onClick={updateTeacher} disabled={loading} type="button"><Save size={18} /> Update</button>
          </div>
        </div>

        <div className="card">
          <h2>Check availability</h2>
          <p>Search uses teacher name. If empty, it uses the name from the form.</p>
          <label className="field">
            <span>Search name</span>
            <input placeholder="Teacher name" value={searchName} onChange={(event) => setSearchName(event.target.value)} />
          </label>
          <div className="actions column">
            <button onClick={readTeacher} disabled={loading} type="button"><Search size={18} /> Read</button>
            <button className="danger" onClick={deleteTeacher} disabled={loading} type="button"><Trash2 size={18} /> Delete</button>
          </div>
        </div>

        <div className="card result-card">
          <div className="card-header">
            <div>
              <h2>Current result</h2>
              <p>Backend returns current-hour availability.</p>
            </div>
            {loading && <RefreshCw className="spin" size={20} />}
          </div>

          {!result ? (
            <div className="empty">No teacher loaded yet.</div>
          ) : (
            <div className="result">
              <div className={`status ${result.isFree ? 'free' : 'busy'}`}>
                {result.isFree ? 'FREE NOW' : 'BUSY NOW'}
              </div>
              <h3>{result.name}</h3>
              {result.free ? (
                <div className="details">
                  <span>Room: <b>{result.free.roomNumber}</b></span>
                  <span>Time: <b>{result.free.from}:00 - {result.free.to}:00</b></span>
                </div>
              ) : (
                <p className="muted">No active free slot for current hour.</p>
              )}
              <pre>{JSON.stringify(result, null, 2)}</pre>
            </div>
          )}
        </div>
      </section>
    </main>
  );
}

createRoot(document.getElementById('root')).render(<App />);
