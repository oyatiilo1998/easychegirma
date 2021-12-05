import 'package:cached_network_image/cached_network_image.dart';
import 'package:chegirma/controller/category_controller.dart';
import 'package:chegirma/core/app_theme.dart';
import 'package:chegirma/core/modal_progress_hud.dart';
import 'package:chegirma/data/model/categories_response.dart';
import 'package:chegirma/ui/product_detail_screen.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:money_formatter/money_formatter.dart';

class CategoryScreen extends GetView<CategoryController> {
  const CategoryScreen(this.categories, {Key? key}) : super(key: key);
  final Categories categories;

  @override
  Widget build(BuildContext context) {
    final width = (Get.width - 52) / 2;
    return GetBuilder<CategoryController>(initState: (_) {
      controller.getProducts(categories.id ?? '');
    }, builder: (logic) {
      return Scaffold(
        appBar: AppBar(
          backgroundColor: Colors.white,
          centerTitle: true,
          elevation: 2,
          title: Text(
            categories.name ?? '',
            style: const TextStyle(
              fontSize: 17,
              fontWeight: FontWeight.bold,
              color: first,
            ),
          ),
        ),
        body: ModalProgressHUD(
          inAsyncCall: logic.loading,
          child: GridView.count(
            crossAxisCount: 2,
            mainAxisSpacing: 20,
            crossAxisSpacing: 16,
            padding: EdgeInsets.all(16),
            childAspectRatio: (width) / (width + 72),
            children: List.generate(
                logic.products.length,
                (index) => InkWell(
                      onTap: () {
                        Get.to(ProductDetailScreen(products: logic.products[index],));
                      },
                      borderRadius: BorderRadius.circular(12),
                      child: Ink(
                        decoration: BoxDecoration(
                          color: Color(0xFFf4f4f4),
                          borderRadius: BorderRadius.circular(12),
                        ),
                        height: width + 72,
                        padding: EdgeInsets.all(4),
                        width: width,
                        child: Column(
                          crossAxisAlignment: CrossAxisAlignment.start,
                          children: [
                            Stack(
                              children: [
                                ClipRRect(
                                  borderRadius: BorderRadius.circular(10),
                                  child: CachedNetworkImage(
                                    imageUrl: logic.products[index].imageUrl ?? '',
                                    height: width,
                                    width: width,
                                    alignment: Alignment.bottomLeft,
                                    fit: BoxFit.cover,
                                    progressIndicatorBuilder: (_, __, ___) {
                                      return SizedBox();
                                    },
                                  ),
                                ),
                                Positioned(
                                  left: 0,
                                  top: 8,
                                  child: Container(
                                    decoration: const BoxDecoration(
                                      color: first,
                                      borderRadius: BorderRadius.horizontal(right: Radius.circular(8))
                                    ),
                                    padding: EdgeInsets.symmetric(horizontal: 6, vertical: 3),
                                    child: Text("-${logic.products[index].discountAmount} %",
                                      style: const TextStyle(
                                        color: Colors.white,
                                        fontSize: 15,
                                        fontWeight: FontWeight.w500,
                                      ),

                                    ),
                                  ),
                                )
                              ],
                            ),
                            Spacer(flex: 2,),
                            Text(
                              logic.products[index].name ?? '',
                              style: const TextStyle(
                                color: Colors.black,
                                fontSize: 15,
                                fontWeight: FontWeight.w500,
                              ),
                            ),
                            Spacer(),
                            Text(
                              moneyFormatWithSymbol(logic.products[index].priceBefore),
                              style: const TextStyle(
                                color: Colors.black,
                                fontSize: 13,
                                decoration: TextDecoration.lineThrough,
                              ),
                            ),
                            Spacer(),
                            Text(
                              moneyFormatWithSymbol(logic.products[index].priceAfter),
                              style: const TextStyle(
                                color: second,
                                fontSize: 15,
                                fontWeight: FontWeight.w500,
                              ),
                            ),
                          ],
                        ),
                      ),
                    )),
          ),
        ),
      );
    });
  }

  static String moneyFormatWithSymbol(int? sum, {String currency = 'USZ', bool abs = true, String separator = ' '}) {
    var amount = MoneyFormatter(
      amount: abs ? (sum ?? 0).roundToDouble().abs() : (sum ?? 0).roundToDouble(),
      settings: MoneyFormatterSettings(
        thousandSeparator: separator,
        symbol: currency,
      ),
    ).output.symbolOnRight;
    return amount.replaceAll('.00', '');
  }
}


