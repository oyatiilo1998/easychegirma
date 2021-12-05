import 'package:chegirma/data/model/product_detail_response.dart';
import 'package:chegirma/data/repository/main_repository.dart';
import 'package:get/get.dart';

class ProductController extends GetxController {
  bool loading = false;
  ProductDetailResponse? response;

  getProducts(String id) async {
    loading = true;
    final result = await MainRepository().getProductDetail(productId: id);
    loading = false;

    if (result.data is ProductDetailResponse) {
      response = result.data;
      update();
    } else {
      Get.snackbar("Error", result.toString());
    }
  }
}
