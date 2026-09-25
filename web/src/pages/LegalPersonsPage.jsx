import { useState } from 'react'

const emptyForm = {
  name: '',
  shortname: '',
  registered_at: '',
  local_kwed: '',
  local_edrpou: '',
}

const formatDate = (value) => {
  if (!value) return '—'
  const date = new Date(value)
  if (Number.isNaN(date.getTime())) return value
  return date.toLocaleDateString('uk-UA')
}

function LegalPersonsPage() {
  const [items, setItems] = useState([])
  const [form, setForm] = useState(emptyForm)
  const [getId, setGetId] = useState('')
  const [selected, setSelected] = useState(null)
  const [deleteId, setDeleteId] = useState('')

  const handleCreate = async (e) => {
    e.preventDefault()

    const payload = {
      ...form,
      registered_at: form.registered_at ? new Date(form.registered_at).toISOString() : '',
      local_edrpou: form.local_edrpou ? Number(form.local_edrpou) : 0,
    }

    const res = await fetch('/legal-persons', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payload),
    })

    const created = await res.json()
    setItems((current) => [created, ...current.filter((it) => it.ID !== created.ID)])
    setForm(emptyForm)
  }

  const handleGet = async () => {
    if (!getId) return
    const res = await fetch(`/legal-persons/${getId}`)
    if (!res.ok) {
      setSelected(null)
      return
    }
    const data = await res.json()
    setSelected(data)
  }

  const handleDelete = async () => {
    if (!deleteId) return
    await fetch(`/legal-persons/${deleteId}`, { method: 'DELETE' })
    setItems((current) => current.filter((it) => String(it.ID) !== String(deleteId)))
    setDeleteId('')
    if (selected && String(selected.ID) === String(deleteId)) {
      setSelected(null)
    }
  }

  return (
    <div className="page-card">
      <div className="page-header">
        <h1>Legal Persons</h1>
      </div>

      <div className="grid-2">
        <section className="panel">
          <h3>Create legal person</h3>
          <form className="form-grid" onSubmit={handleCreate}>
            <input value={form.name} onChange={(e) => setForm({ ...form, name: e.target.value })} placeholder="Name" required />
            <input value={form.shortname} onChange={(e) => setForm({ ...form, shortname: e.target.value })} placeholder="Short name" />
            <input type="date" value={form.registered_at} onChange={(e) => setForm({ ...form, registered_at: e.target.value })} />
            <input value={form.local_kwed} onChange={(e) => setForm({ ...form, local_kwed: e.target.value })} placeholder="Local KWED" />
            <input value={form.local_edrpou} onChange={(e) => setForm({ ...form, local_edrpou: e.target.value })} placeholder="Local ERDPOU" />
            <button className="primary-btn" type="submit">Create</button>
          </form>
        </section>

        <section className="panel">
          <h3>Get legal person by ID</h3>
          <div className="form-grid">
            <input value={getId} onChange={(e) => setGetId(e.target.value)} placeholder="Legal person ID" />
            <button className="secondary-btn" onClick={handleGet}>Get</button>
            {selected ? (
              <div className="item-main">
                <span className="tag">ID {selected.ID}</span>
                <div className="item-title">{selected.Name}</div>
                <div className="item-meta">Short: {selected.ShortName || '—'}</div>
                <div className="item-meta">Registered: {formatDate(selected.RegisteredAt)}</div>
                <div className="item-meta">KWED: {selected.LocalKWED || '—'} • ERDPOU: {selected.LocalERDPOU || '—'}</div>
              </div>
            ) : (
              <div className="empty">No legal person selected</div>
            )}
          </div>
        </section>
      </div>

      <section className="panel" style={{ marginTop: 20 }}>
        <h3>Delete legal person by ID</h3>
        <div className="form-grid">
          <input value={deleteId} onChange={(e) => setDeleteId(e.target.value)} placeholder="Legal person ID" />
          <button className="danger-btn" onClick={handleDelete}>Delete</button>
        </div>
      </section>

      {items.length > 0 && (
        <section className="panel" style={{ marginTop: 20 }}>
          <h3>Recent legal persons</h3>
          <ul className="list">
            {items.map((item) => (
              <li key={item.ID} className="item">
                <div className="item-main">
                  <span className="item-title">{item.Name}</span>
                  <span className="item-meta">ID: {item.ID} • Short: {item.ShortName || '—'} • ERDPOU: {item.LocalERDPOU || '—'}</span>
                </div>
              </li>
            ))}
          </ul>
        </section>
      )}
    </div>
  )
}

export default LegalPersonsPage
