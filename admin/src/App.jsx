import AppRouter from "./router"
import { Router } from "react-router-dom"
import history from "./history"

function App() {

  return (
    <>
      <Router history={history} >
        <AppRouter />
      </Router>
    </>
  )
}

export default App
