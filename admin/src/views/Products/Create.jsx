import { Button, Card, experimentalStyled, Typography } from "@mui/material"
import { useFormik } from "formik"
import { useEffect, useState } from "react"
import { useSelector } from "react-redux"
import { useHistory, useLocation, useParams } from "react-router"
import FDatePicker from "../../components/FormElements/FDatePicker"
import FImageUpload from "../../components/FormElements/FImageUpload"
import FRow from "../../components/FormElements/FRow"
import FSelect from "../../components/FormElements/FSelect"
import FTextField from "../../components/FormElements/FTextField2"
import FullPageLoader from "../../components/FullPageLoader"
import Header from "../../components/Header"
import categoryService from "../../services/categoryService"
import productService from "../../services/productService"

const StyledCard = experimentalStyled(Card)(({ theme }) => ({
  display: "block",
  padding: 15,
  marginBottom: 20,
  maxWidth: "700px",
}))

const initialValues = {
  category_id: null,
  discount_amount: "",
  image_url: "",
  name: "",
  price_after: "",
  price_before: "",
  url: "",
  from_time: null,
  to_time: null,
  description: "",
}

const ProductCreate = () => {
  const history = useHistory()
  const params = useParams()

  const [loader, setLoader] = useState(true)
  const [btnLoader, setBtnLoader] = useState(false)
  const [categoriesList, setCategoriesList] = useState([])
  const userId = useSelector((state) => state.auth.userId)

  useEffect(() => {
    fetchCategoriesList()
    fetchData()
  }, [])

  const fetchData = () => {
    if (!params.id) return setLoader(false)

    productService
      .getProductById(params.id)
      .then((res) => {
        formik.setFieldValue("category_id", res.category.id)
        formik.setFieldValue("discount_amount", res.discount_amount)
        formik.setFieldValue("image_url", res.image_url)
        formik.setFieldValue("price_after", res.price_after)
        formik.setFieldValue("price_before", res.price_before)
        formik.setFieldValue("url", res.url)
        formik.setFieldValue("name", res.name)
        formik.setFieldValue("from_time", res.from_time)
        formik.setFieldValue("to_time", res.to_time)
        formik.setFieldValue("description", res.description)
      })
      .finally(() => setLoader(false))
  }

  const fetchCategoriesList = () => {
    categoryService.getCategoryList().then((res) =>
      setCategoriesList(
        res?.categories.map((category) => ({
          label: category.name,
          value: category.id,
        }))
      )
    )
  }

  const onSubmit = (values) => {
    const computedValues = {
      ...values,
      discount_amount: Number(values.discount_amount),
      price_after: Number(values.price_after),
      price_before: Number(values.price_before),
    }

    if (params.id) return updateMenuElementData(computedValues)
    createMenuElement(computedValues)
  }

  const createMenuElement = (data) => {
    setBtnLoader(true)
    productService
      .createProduct(data)
      .then((res) => {
        history.push("/main/products")
      })
      .finally(() => setBtnLoader(false))
  }

  const updateMenuElementData = (data) => {
    setBtnLoader(true)
    productService
      .updateProduct(params.id, data)
      .then((res) => {
        history.push("/main/products")
      })
      .finally(() => setBtnLoader(false))
  }

  const formik = useFormik({
    initialValues: {
      ...initialValues,
      owner_id: userId,
    },
    onSubmit,
  })

  if (loader) return <FullPageLoader />

  return (
    <form onSubmit={formik.handleSubmit}>
      <Header title="Продукты">
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

          <FRow label="Категория">
            <FSelect
              options={categoriesList}
              width={"100%"}
              formik={formik}
              name="category_id"
            />
          </FRow>

          <FRow label="Фото">
            <FImageUpload formik={formik} name="image_url" />
          </FRow>

          <FRow label="Описание">
            <FTextField fullWidth multiline rows={4} formik={formik} name="description" />
          </FRow>

          <FRow label="Скидка">
            <FTextField fullWidth formik={formik} name="discount_amount" />
          </FRow>

          <FRow label="Цена до скидки">
            <FTextField fullWidth formik={formik} name="price_before" />
          </FRow>

          <FRow label="Цена после скидки">
            <FTextField fullWidth formik={formik} name="price_after" />
          </FRow>

          <FRow label="Ссылка на товар">
            <FTextField fullWidth formik={formik} name="url" />
          </FRow>

          <FRow label="Дата начала акции">
            <FDatePicker width="100%" formik={formik} name="from_time" />
          </FRow>

          <FRow label="Дата окончания акции">
            <FDatePicker width="100%" formik={formik} name="to_time" />
          </FRow>
        </StyledCard>
      </div>
    </form>
  )
}

export default ProductCreate
