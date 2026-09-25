import { useState } from 'react'

const emptyPersonFilter = {
  contract_type: '',
  start_date_from: '',
  start_date_to: '',
  person_name: '',
  person_lastname: '',
  person_patronym: '',
  person_date_of_birth: '',
}

const emptyLegalFilter = {
  contract_type: '',
  start_date_from: '',
  start_date_to: '',
  legal_shortname: '',
  legal_name: '',
  legal_edrpou: '',
}

const formatDate = (value) => {
  if (!value) return '—'
  const date = new Date(value)
  if (Number.isNaN(date.getTime())) return value
  return date.toLocaleDateString('uk-UA')
}

const buildPayload = (source) => {
  const payload = {}

  Object.entries(source).forEach(([key, value]) => {
    if (value === '' || value === null || value === undefined) return

    if (key.endsWith('_from') || key.endsWith('_to') || key.endsWith('_date_of_birth')) {
      payload[key] = new Date(value).toISOString()
      return
    }

    if (key === 'legal_edrpou') {
      const num = Number(value)
      if (Number.isFinite(num) && num > 0) payload[key] = num
      return
    }

    payload[key] = value
  })

  return payload
}

function SearchPage() {
  const [personFilter, setPersonFilter] = useState(emptyPersonFilter)
  const [legalFilter, setLegalFilter] = useState(emptyLegalFilter)
  const [personResults, setPersonResults] = useState([])
  const [legalResults, setLegalResults] = useState([])

  const handlePersonSearch = async (e) => {
    e.preventDefault()

    const payload = buildPayload(personFilter)

    const res = await fetch('/persons/search', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payload),
    })

    const data = await res.json()
    setPersonResults(Array.isArray(data) ? data : [])
  }

  const handleLegalSearch = async (e) => {
    e.preventDefault()

    const payload = buildPayload(legalFilter)

    const res = await fetch('/legal-persons/search', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payload),
    })

    const data = await res.json()
    setLegalResults(Array.isArray(data) ? data : [])
  }

  return (
    <div className="page-card">
      <div className="page-header">
        <h1>Contract Search</h1>
      </div>

      <div className="grid-2">
        <section className="panel">
          <h3>Search contracts by person</h3>
          <form className="form-grid" onSubmit={handlePersonSearch}>
            <input value={personFilter.contract_type} onChange={(e) => setPersonFilter({ ...personFilter, contract_type: e.target.value })} placeholder="Contract type" />
            <input type="date" value={personFilter.start_date_from} onChange={(e) => setPersonFilter({ ...personFilter, start_date_from: e.target.value })} />
            <input type="date" value={personFilter.start_date_to} onChange={(e) => setPersonFilter({ ...personFilter, start_date_to: e.target.value })} />
            <input value={personFilter.person_name} onChange={(e) => setPersonFilter({ ...personFilter, person_name: e.target.value })} placeholder="Person name" />
            <input value={personFilter.person_lastname} onChange={(e) => setPersonFilter({ ...personFilter, person_lastname: e.target.value })} placeholder="Person lastname" />
            <input value={personFilter.person_patronym} onChange={(e) => setPersonFilter({ ...personFilter, person_patronym: e.target.value })} placeholder="Person patronym" />
            <input type="date" value={personFilter.person_date_of_birth} onChange={(e) => setPersonFilter({ ...personFilter, person_date_of_birth: e.target.value })} />
            <button className="primary-btn" type="submit">Search</button>
          </form>

          <div style={{ marginTop: 16 }}>
            {personResults.length === 0 ? (
              <div className="empty">No matches</div>
            ) : (
              <ul className="list">
                {personResults.map((contract) => (
                  <li key={contract.ID} className="item">
                    <div className="item-main">
                      <span className="item-title">#{contract.Number} • {contract.Type || 'No type'}</span>
                      <span className="item-meta">ID: {contract.ID} • {formatDate(contract.StartDate)} – {formatDate(contract.EndDate)}</span>
                    </div>
                  </li>
                ))}
              </ul>
            )}
          </div>
        </section>

        <section className="panel">
          <h3>Search contracts by legal person</h3>
          <form className="form-grid" onSubmit={handleLegalSearch}>
            <input value={legalFilter.contract_type} onChange={(e) => setLegalFilter({ ...legalFilter, contract_type: e.target.value })} placeholder="Contract type" />
            <input type="date" value={legalFilter.start_date_from} onChange={(e) => setLegalFilter({ ...legalFilter, start_date_from: e.target.value })} />
            <input type="date" value={legalFilter.start_date_to} onChange={(e) => setLegalFilter({ ...legalFilter, start_date_to: e.target.value })} />
            <input value={legalFilter.legal_shortname} onChange={(e) => setLegalFilter({ ...legalFilter, legal_shortname: e.target.value })} placeholder="Legal shortname" />
            <input value={legalFilter.legal_name} onChange={(e) => setLegalFilter({ ...legalFilter, legal_name: e.target.value })} placeholder="Legal name" />
            <input value={legalFilter.legal_edrpou} onChange={(e) => setLegalFilter({ ...legalFilter, legal_edrpou: e.target.value })} placeholder="Legal EDRPOU" />
            <button className="primary-btn" type="submit">Search</button>
          </form>

          <div style={{ marginTop: 16 }}>
            {legalResults.length === 0 ? (
              <div className="empty">No matches</div>
            ) : (
              <ul className="list">
                {legalResults.map((contract) => (
                  <li key={contract.ID} className="item">
                    <div className="item-main">
                      <span className="item-title">#{contract.Number} • {contract.Type || 'No type'}</span>
                      <span className="item-meta">ID: {contract.ID} • {formatDate(contract.StartDate)} – {formatDate(contract.EndDate)}</span>
                    </div>
                  </li>
                ))}
              </ul>
            )}
          </div>
        </section>
      </div>
    </div>
  )
}

export default SearchPage
