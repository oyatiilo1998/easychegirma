import { useHistory } from "react-router"
import "./style.scss"

const CatalogCard = ({category}) => {
  const history = useHistory()

  return (
    <div className="CatalogCard" onClick={() => history.push(`/category/${category?.id}`)} >
      <img src={category?.ru_name} alt="" className="catalog-image" />

      <h3 className="catalog-name">{category?.name}</h3>
    </div>
  )
}

export default CatalogCard
