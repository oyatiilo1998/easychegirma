import CatalogCard from "../../../components/CatalogCard"
import "./style.scss"

const CatalogsBlock = ({categoriesList}) => {
  return (
    <>
      <h2 className="main-title">Категории скидок</h2>
      <div className="CatalogsBlock" >

        {
          categoriesList.map(category => (
            <CatalogCard category={category} />
          ))
        }
        {/* <CatalogCard /> */}
        
      </div>
    </>
  )
}

export default CatalogsBlock
