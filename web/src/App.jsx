import { BrowserRouter, NavLink, Route, Routes } from 'react-router-dom'
import PersonsPage from './pages/PersonsPage'
import LegalPersonsPage from './pages/LegalPersonsPage'
import ContractsPage from './pages/ContractsPage'
import SearchPage from './pages/SearchPage'

function App() {
  return (
    <BrowserRouter>
      <div className="app-shell">
        <header className="topbar">
          <div className="brand">Contracts CRM</div>
          <nav className="navbar" aria-label="Main navigation">
            <NavLink to="/persons" className={({ isActive }) => 'nav-link ' + (isActive ? 'active' : '')}>
              Persons
            </NavLink>
            <NavLink to="/legal-persons" className={({ isActive }) => 'nav-link ' + (isActive ? 'active' : '')}>
              Legal Persons
            </NavLink>
            <NavLink to="/contracts" className={({ isActive }) => 'nav-link ' + (isActive ? 'active' : '')}>
              Contracts
            </NavLink>
            <NavLink to="/search" className={({ isActive }) => 'nav-link ' + (isActive ? 'active' : '')}>
              Search
            </NavLink>
          </nav>
        </header>

        <main className="page-shell">
          <Routes>
            <Route path="/" element={<PersonsPage />} />
            <Route path="/persons" element={<PersonsPage />} />
            <Route path="/legal-persons" element={<LegalPersonsPage />} />
            <Route path="/contracts" element={<ContractsPage />} />
            <Route path="/search" element={<SearchPage />} />
          </Routes>
        </main>
      </div>
    </BrowserRouter>
  )
}

export default App
