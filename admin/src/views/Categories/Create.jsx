import { Button, Card, experimentalStyled, Typography } from "@mui/material"
import { useFormik } from "formik"
import { useEffect, useState } from "react"
import { useHistory, useLocation, useParams } from "react-router"
import FImageUpload from "../../components/FormElements/FImageUpload"
import FRow from "../../components/FormElements/FRow"
import FTextField from "../../components/FormElements/FTextField2"
import FullPageLoader from "../../components/FullPageLoader"
import Header from "../../components/Header"
import categoryService from "../../services/categoryService"

const StyledCard = experimentalStyled(Card)(({ theme }) => ({
  display: "block",
  padding: 15,
  marginBottom: 20,
  maxWidth: "700px",
}))

const initialValues = {
  name: "",
  ru_name: "",
  code: 1
}

const CategoryCreate = () => {
  const history = useHistory()
  const params = useParams()

  const [loader, setLoader] = useState(true)
  const [btnLoader, setBtnLoader] = useState(false)


  useEffect(() => {
    fetchData()
  }, [])

  const fetchData = () => {
    if (!params.id) return setLoader(false)

    categoryService
      .getCategoryById(params.id)
      .then((res) => {
          formik.setFieldValue("name", res.name)
          formik.setFieldValue("ru_name", res.ru_name)
      })
      .finally(() => setLoader(false))
  }

  const onSubmit = (values) => {
    if (params.id) return updateMenuElementData(values)
    createMenuElement(values)
  }

  const createMenuElement = (data) => {
    setBtnLoader(true)
    categoryService
      .createCategory(data)
      .then((res) => {
        history.push("/main/categories")
      })
      .finally(() => setBtnLoader(false))
  }

  const updateMenuElementData = (data) => {
    setBtnLoader(true)
    categoryService
      .updateCategory(params.id, data)
      .then((res) => {
        history.push("/main/categories")
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
      <Header title="Категории">
        <Button variant="contained" color="error">
          Отменить
        </Button>
        <Button disabled={btnLoader} variant="contained" type="submit">
          Сохранить
        </Button>
      </Header>

      <div className="form-container">
        <StyledCard elevation={12}>
          <FRow label="Название">
            <FTextField fullWidth formik={formik} name="name" />
          </FRow>
          <FRow label="Фото">
            <FImageUpload formik={formik} name="ru_name" />
          </FRow>
        </StyledCard>
      </div>
    </form>
  )
}

export default CategoryCreate
