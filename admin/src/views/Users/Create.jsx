import { Button, Card, experimentalStyled, Typography } from "@mui/material"
import { useFormik } from "formik"
import { useEffect, useState } from "react"
import { useHistory, useLocation, useParams } from "react-router"
import FImageUpload from "../../components/FormElements/FImageUpload"
import FRow from "../../components/FormElements/FRow"
import FSelect from "../../components/FormElements/FSelect"
import FTextField from "../../components/FormElements/FTextField2"
import FullPageLoader from "../../components/FullPageLoader"
import Header from "../../components/Header"
import categoryService from "../../services/categoryService"
import productService from "../../services/productService"
import userService from "../../services/userService"

const StyledCard = experimentalStyled(Card)(({ theme }) => ({
  display: "block",
  padding: 15,
  marginBottom: 20,
  maxWidth: "700px",
}))

const initialValues = {
  image_url: "",
  link: "",
  login: "",
  name: "",
  password: "",
}

const UserCreate = () => {
  const history = useHistory()
  const params = useParams()

  const [loader, setLoader] = useState(true)
  const [btnLoader, setBtnLoader] = useState(false)

  useEffect(() => {
    fetchData()
  }, [])

  const fetchData = () => {
    if (!params.id) return setLoader(false)

    userService
      .getUserById(params.id)
      .then((res) => {
        formik.setFieldValue("image_url", res.image_url)
        formik.setFieldValue("link", res.link)
        formik.setFieldValue("login", res.login)
        formik.setFieldValue("name", res.name)
        formik.setFieldValue("password", res.password)
      })
      .finally(() => setLoader(false))
  }

  const onSubmit = (values) => {
    if (params.id) return updateMenuElementData(values)
    createMenuElement(values)
  }

  const createMenuElement = (data) => {
    setBtnLoader(true)
    userService
      .createUser(data)
      .then((res) => {
        history.push("/main/users")
      })
      .finally(() => setBtnLoader(false))
  }

  const updateMenuElementData = (data) => {
    setBtnLoader(true)
    userService
      .updateUser(params.id, data)
      .then((res) => {
        history.push("/main/users")
      })
      .finally(() => setBtnLoader(false))
  }

  const formik = useFormik({
    initialValues,
    onSubmit,
  })

  if (loader) return <FullPageLoader />

  return (
    <form onSubmit={formik.handleSubmit}>
      <Header title="Пользователи">
        <Button variant="contained" color="error">
          Отменить
        </Button>
        <Button disabled={btnLoader} variant="contained" type="submit">
          Сохранить
        </Button>
      </Header>

      <div className="form-container">
        <StyledCard elevation={12}>
          <FRow label="Имя">
            <FTextField fullWidth formik={formik} name="name" />
          </FRow>

          <FRow label="Фото">
            <FImageUpload formik={formik} name="image_url" />
          </FRow>

          <FRow label="Логин">
            <FTextField fullWidth formik={formik} name="login" />
          </FRow>

          <FRow label="Пароль">
            <FTextField type="password" fullWidth formik={formik} name="password" />
          </FRow>

          <FRow label="Линк">
            <FTextField fullWidth formik={formik} name="link" />
          </FRow>

        </StyledCard>
      </div>
    </form>
  )
}

export default UserCreate
