import { useState } from 'react'

const emptyForm = {
  number: '',
  start_date: '',
  end_date: '',
  type: '',
}

const formatDate = (value) => {
  if (!value) return '—'
  const date = new Date(value)
  if (Number.isNaN(date.getTime())) return value
  return date.toLocaleDateString('uk-UA')
}

function ContractsPage() {
  const [items, setItems] = useState([])
  const [form, setForm] = useState(emptyForm)
  const [getId, setGetId] = useState('')
  const [selected, setSelected] = useState(null)
  const [deleteId, setDeleteId] = useState('')
  const [relationPerson, setRelationPerson] = useState({ contract_id: '', person_id: '' })
  const [relationLegalPerson, setRelationLegalPerson] = useState({ contract_id: '', legal_person_id: '' })

  const handleCreate = async (e) => {
    e.preventDefault()
    const payload = {
      ...form,
      number: Number(form.number),
      start_date: form.start_date ? new Date(form.start_date).toISOString() : '',
      end_date: form.end_date ? new Date(form.end_date).toISOString() : '',
    }

    const res = await fetch('/contracts', {
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
    const res = await fetch(`/contracts/${getId}`)
    if (!res.ok) {
      setSelected(null)
      return
    }
    const data = await res.json()
    setSelected(data)
  }

  const handleDelete = async () => {
    if (!deleteId) return
    await fetch(`/contracts/${deleteId}`, { method: 'DELETE' })
    setItems((current) => current.filter((it) => String(it.ID) !== String(deleteId)))
    setDeleteId('')
    if (selected && selected.contract && String(selected.contract.ID) === String(deleteId)) {
      setSelected(null)
    }
  }

  const handleAddPerson = async () => {
    if (!relationPerson.contract_id || !relationPerson.person_id) return
    await fetch(`/contracts/${relationPerson.contract_id}/persons`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ person_id: Number(relationPerson.person_id) }),
    })
    setRelationPerson({ contract_id: '', person_id: '' })
    if (getId) {
      await handleGet()
    }
  }

  const handleAddLegalPerson = async () => {
    if (!relationLegalPerson.contract_id || !relationLegalPerson.legal_person_id) return
    await fetch(`/contracts/${relationLegalPerson.contract_id}/legal-persons`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ legal_person_id: Number(relationLegalPerson.legal_person_id) }),
    })
    setRelationLegalPerson({ contract_id: '', legal_person_id: '' })
    if (getId) {
      await handleGet()
    }
  }

  const handleRemovePerson = async (contractId, personId) => {
    await fetch(`/contracts/${contractId}/persons/${personId}`, { method: 'DELETE' })
    await handleGet()
  }

  const handleRemoveLegalPerson = async (contractId, legalPersonId) => {
    await fetch(`/contracts/${contractId}/legal-persons/${legalPersonId}`, { method: 'DELETE' })
    await handleGet()
  }

  return (
    <div className="page-card">
      <div className="page-header">
        <h1>Contracts</h1>
      </div>

      <div className="grid-2">
        <section className="panel">
          <h3>Create contract</h3>
          <form className="form-grid" onSubmit={handleCreate}>
            <input type="number" value={form.number} onChange={(e) => setForm({ ...form, number: e.target.value })} placeholder="Number" required />
            <input type="date" value={form.start_date} onChange={(e) => setForm({ ...form, start_date: e.target.value })} />
            <input type="date" value={form.end_date} onChange={(e) => setForm({ ...form, end_date: e.target.value })} />
            <input value={form.type} onChange={(e) => setForm({ ...form, type: e.target.value })} placeholder="Type" />
            <button className="primary-btn" type="submit">Create</button>
          </form>
        </section>

        <section className="panel">
          <h3>Get contract by ID</h3>
          <div className="form-grid">
            <input value={getId} onChange={(e) => setGetId(e.target.value)} placeholder="Contract ID" />
            <button className="secondary-btn" onClick={handleGet}>Get</button>
            {selected ? (
              <div className="item-main">
                <span className="tag">ID {selected.contract?.id}</span>
                <div className="item-title">#{selected.contract?.number}</div>
                <div className="item-meta">{selected.contract?.type || '—'}</div>
                <div className="item-meta">From {formatDate(selected.contract?.start_date)} to {formatDate(selected.contract?.end_date)}</div>
                <div className="item-meta">Persons: {selected.persons?.length || 0} • Legal persons: {selected.legal_persons?.length || 0}</div>
              </div>
            ) : (
              <div className="empty">No contract selected</div>
            )}
          </div>
        </section>
      </div>

      <div className="grid-2" style={{ marginTop: 20 }}>
        <section className="panel">
          <h3>Add person to contract</h3>
          <div className="form-grid">
            <input value={relationPerson.contract_id} onChange={(e) => setRelationPerson({ ...relationPerson, contract_id: e.target.value })} placeholder="Contract ID" />
            <input value={relationPerson.person_id} onChange={(e) => setRelationPerson({ ...relationPerson, person_id: e.target.value })} placeholder="Person ID" />
            <button className="primary-btn" onClick={handleAddPerson}>Add person</button>
          </div>
        </section>

        <section className="panel">
          <h3>Add legal person to contract</h3>
          <div className="form-grid">
            <input value={relationLegalPerson.contract_id} onChange={(e) => setRelationLegalPerson({ ...relationLegalPerson, contract_id: e.target.value })} placeholder="Contract ID" />
            <input value={relationLegalPerson.legal_person_id} onChange={(e) => setRelationLegalPerson({ ...relationLegalPerson, legal_person_id: e.target.value })} placeholder="Legal person ID" />
            <button className="primary-btn" onClick={handleAddLegalPerson}>Add legal person</button>
          </div>
        </section>
      </div>

      <section className="panel" style={{ marginTop: 20 }}>
        <h3>Delete contract by ID</h3>
        <div className="form-grid">
          <input value={deleteId} onChange={(e) => setDeleteId(e.target.value)} placeholder="Contract ID" />
          <button className="danger-btn" onClick={handleDelete}>Delete</button>
        </div>
      </section>

      {items.length > 0 && (
        <section className="panel" style={{ marginTop: 20 }}>
          <h3>Recent contracts</h3>
          <ul className="list">
            {items.map((contract) => (
                  <li key={contract.id} className="item">
                    <div className="item-main">
                      <span className="item-title">
                        #{contract.number} • {contract.type || 'No type'}
                      </span>
                      <span className="item-meta">
                        ID: {contract.id} • {formatDate(contract.start_date)} – {formatDate(contract.end_date)}
                      </span>
                    </div>
                  </li>
            ))}
          </ul>
        </section>
      )}

      {selected && (
        <section className="panel" style={{ marginTop: 20 }}>
          <h3>Details for contract #{selected.contract?.number}</h3>
          <div className="grid-2">
            <div>
              <h4>Persons</h4>
              {selected.persons?.length ? (
                <ul className="list">
                  {selected.persons.map((person) => (
                    <li key={person.ID} className="item">
                      <div className="item-main">
                        <span className="item-title">{person.lastname} {person.name} {person.patronym}</span>
                        <span className="item-meta">ID: {person.id}</span>
                      </div>
                      <button className="danger-btn" onClick={() => handleRemovePerson(selected.contract.id, person.id)}>Remove</button>
                    </li>
                  ))}
                </ul>
              ) : (
                <div className="empty">No persons linked</div>
              )}
            </div>

            <div>
              <h4>Legal persons</h4>
              {selected.legalPersons?.length ? (
                <ul className="list">
                  {selected.legalPersons.map((lp) => (
                    <li key={lp.ID} className="item">
                      <div className="item-main">
                        <span className="item-title">{lp.name}</span>
                        <span className="item-meta">ID: {lp.id}</span>
                      </div>
                      <button className="danger-btn" onClick={() => handleRemoveLegalPerson(selected.contract.id, lp.id)}>Remove</button>
                    </li>
                  ))}
                </ul>
              ) : (
                <div className="empty">No legal persons linked</div>
              )}
            </div>
          </div>
        </section>
      )}
    </div>
  )
}

export default ContractsPage
