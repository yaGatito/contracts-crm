import { useState } from 'react'

const emptyForm = {
  lastname: '',
  name: '',
  patronym: '',
  date_of_birth: '',
}

const formatDate = (value) => {
  if (!value) return '—'
  const date = new Date(value)
  if (Number.isNaN(date.getTime())) return value
  return date.toLocaleDateString('uk-UA')
}

function PersonsPage() {
  const [items, setItems] = useState([])
  const [form, setForm] = useState(emptyForm)
  const [getId, setGetId] = useState('')
  const [selected, setSelected] = useState(null)
  const [deleteId, setDeleteId] = useState('')

  const handleCreate = async (e) => {
    e.preventDefault()
    const payload = {
      ...form,
      date_of_birth: form.date_of_birth ? new Date(form.date_of_birth).toISOString() : '',
    }

    const res = await fetch('/persons', {
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
    const res = await fetch(`/persons/${getId}`)
    if (!res.ok) {
      setSelected(null)
      return
    }
    const data = await res.json()
    setSelected(data)
  }

  const handleDelete = async () => {
    if (!deleteId) return
    await fetch(`/persons/${deleteId}`, { method: 'DELETE' })
    setItems((current) => current.filter((it) => String(it.ID) !== String(deleteId)))
    setDeleteId('')
    if (selected && String(selected.ID) === String(deleteId)) {
      setSelected(null)
    }
  }

  return (
    <div className="page-card">
      <div className="page-header">
        <h1>Persons</h1>
      </div>

      <div className="grid-2">
        <section className="panel">
          <h3>Create person</h3>
          <form className="form-grid" onSubmit={handleCreate}>
            <input value={form.lastname} onChange={(e) => setForm({ ...form, lastname: e.target.value })} placeholder="Lastname" required />
            <input value={form.name} onChange={(e) => setForm({ ...form, name: e.target.value })} placeholder="Name" required />
            <input value={form.patronym} onChange={(e) => setForm({ ...form, patronym: e.target.value })} placeholder="Patronym" />
            <input type="date" value={form.date_of_birth} onChange={(e) => setForm({ ...form, date_of_birth: e.target.value })} />
            <button className="primary-btn" type="submit">Create</button>
          </form>
        </section>

        <section className="panel">
          <h3>Get person by ID</h3>
          <div className="form-grid">
            <input value={getId} onChange={(e) => setGetId(e.target.value)} placeholder="Person ID" />
            <button className="secondary-btn" onClick={handleGet}>Get</button>
            {selected ? (
              <div className="item-main">
                <span className="tag">ID {selected.ID}</span>
                <div className="item-title">{selected.Lastname} {selected.Name} {selected.Patronym}</div>
                <div className="item-meta">Date of birth: {formatDate(selected.DateOfBirth)}</div>
              </div>
            ) : (
              <div className="empty">No person selected</div>
            )}
          </div>
        </section>
      </div>

      <section className="panel" style={{ marginTop: 20 }}>
        <h3>Delete person by ID</h3>
        <div className="form-grid">
          <input value={deleteId} onChange={(e) => setDeleteId(e.target.value)} placeholder="Person ID" />
          <button className="danger-btn" onClick={handleDelete}>Delete</button>
        </div>
      </section>

      {items.length > 0 && (
        <section className="panel" style={{ marginTop: 20 }}>
          <h3>Recent persons</h3>
          <ul className="list">
            {items.map((person) => (
              <li key={person.ID} className="item">
                <div className="item-main">
                  <span className="item-title">{person.Lastname} {person.Name} {person.Patronym}</span>
                  <span className="item-meta">ID: {person.ID} • {formatDate(person.DateOfBirth)}</span>
                </div>
              </li>
            ))}
          </ul>
        </section>
      )}
    </div>
  )
}

export default PersonsPage
