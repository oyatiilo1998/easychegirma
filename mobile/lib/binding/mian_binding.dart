import 'package:chegirma/controller/category_controller.dart';
import 'package:chegirma/controller/home_controller.dart';
import 'package:chegirma/controller/product_controller.dart';
import 'package:get/get.dart';

class MainBinding extends Bindings{

  @override
  void dependencies() {
    Get.lazyPut<HomeController>(() => HomeController(),
      fenix: true,
    );
    Get.lazyPut<CategoryController>(() => CategoryController(),
      fenix: true,
    );
    Get.lazyPut<ProductController>(() => ProductController(),
      fenix: true,
    );
  }
}