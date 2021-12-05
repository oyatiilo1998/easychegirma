import { Router } from "react-router"
import { BrowserRouter } from "react-router-dom"
import AppRouter from "./router"

function App() {
  return (
    <div className="App">
      <BrowserRouter>
        <AppRouter />
      </BrowserRouter>
    </div>
  )
}

export default App