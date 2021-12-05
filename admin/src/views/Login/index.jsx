import { Button, TextField, Typography } from "@mui/material"
import { useState } from "react"
import { useSelector } from "react-redux"
import { useDispatch } from "react-redux"
import Logo from "../../components/Logo"
import { authActions } from "../../redux/slices/auth.slice"
import { showAlert } from "../../redux/thunks/alert.thunk"
import authService from "../../services/authService"
import "./style.scss"

const LoginPage = () => {
  const dispatch = useDispatch()
  const [btnLoader, setBtnLoader] = useState(false)
  const [login, setLogin] = useState("")
  const [password, setPassword] = useState("")

  const onSubmit = (e) => {
    e.preventDefault()
    setBtnLoader(true)
    authService.checkLoginAndPassword({
      login,
      password
    }).then(res => {
      dispatch(authActions.login({
        userId: res.id,
        userLogin: login
      }))
    })
    .catch(err => dispatch(showAlert('Вы ввели неправильный логин или пароль')))
    .finally(() => setBtnLoader(false))
  }

  return (
    <div className="LoginPage">
      <div className="section left">
        <Logo />
      </div>
      <div className="section right">
        <form onSubmit={onSubmit} className="form">
          <Typography
            style={{ marginBottom: "50px" }}
            variant="h2"
            color="secondary"
          >
            Вход в систему
          </Typography>
          <TextField
            value={login}
            onChange={(e) => setLogin(e.target.value)}
            style={{ marginBottom: "30px" }}
            fullWidth
            label="Логин"
          />
          <TextField
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            style={{ marginBottom: "30px" }}
            fullWidth
            label="Пароль"
          />

          <Button
            disabled={btnLoader}
            // onClick={onSubmit}
            type="submit"
            color="secondary"
            variant="contained"
            fullWidth
            size="large"
          >
            Войти
          </Button>
        </form>
      </div>
    </div>
  )
}

export default LoginPage
