import ReactModal from "react-modal"
import Button from "../Button"
import "./style.scss"

const customStyles = {
  overlay: {
    backgroundColor: "rgba(0, 0, 0, .8)",
    zIndex: 101,
  },

  content: {
    top: "50%",
    left: "50%",
    right: "auto",
    bottom: "auto",
    marginRight: "-50%",
    transform: "translate(-50%, -50%)",
    borderRadius: "20px",
    width: "900px",
  },
}

const ProductModal = ({selectedProduct, ...props}) => {
  return (
    <ReactModal
      style={customStyles}
      {...props}
    >
      <div className="ProductModal">
        <div className="left-side">
          <img
            src={selectedProduct?.image_url}
            alt=""
            className="product-image"
          />
        </div>
        <div className="right-side">
          <div className="top">
            <h2 className="product-name">{selectedProduct?.name}</h2>
            <p className="description">{selectedProduct?.description ?? '----'}</p>
            <div className="category-name" style={{ color: 'silver', margin: '5px 0' }} >{selectedProduct?.owner?.name}</div>

            <div className="discount">{selectedProduct?.discount_amount} %</div>
            <div className="price-row">
              <span className="old-price">{selectedProduct?.price_before} UZS</span> -{" "}
              <span className="new-price">{selectedProduct?.price_after} UZS</span>
            </div>
          </div>
          <div className="bottom">
            <Button className="third" style={{flex: 1}} onClick={props.onRequestClose} >Назад</Button>
            <a href={`${selectedProduct?.url}`} target="_blank" rel="noreferrer"><Button className="secondary" style={{flex: 1}} >Перейти в магазин</Button></a>
          </div>
        </div>
      </div>
    </ReactModal>
  )
}

export default ProductModal
