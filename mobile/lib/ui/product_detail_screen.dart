import 'package:cached_network_image/cached_network_image.dart';
import 'package:chegirma/controller/product_controller.dart';
import 'package:chegirma/core/app_theme.dart';
import 'package:chegirma/core/modal_progress_hud.dart';
import 'package:chegirma/data/model/products_response.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:money_formatter/money_formatter.dart';
import 'package:url_launcher/url_launcher.dart';

class ProductDetailScreen extends GetView<ProductController> {
  const ProductDetailScreen({Key? key, required this.products}) : super(key: key);
  final Products products;


  @override
  Widget build(BuildContext context) {
    return GetBuilder<ProductController>(
        initState: (_){
          controller.getProducts(products.id??'');
        },
        builder: (logic) {
      return Scaffold(
        appBar: AppBar(
          backgroundColor: Colors.white,
          centerTitle: true,
          elevation: 2,
          title: Text(
            products.name ?? '',
            style: const TextStyle(
              fontSize: 17,
              fontWeight: FontWeight.bold,
              color: first,
            ),
          ),
        ),
        body: ModalProgressHUD(
          inAsyncCall: logic.loading,
          child: logic.response!=null? ListView(
            children: [
              CachedNetworkImage(
                imageUrl: logic.response!.imageUrl ?? '',
                height: 208,
                width: Get.width,
                alignment: Alignment.bottomLeft,
                fit: BoxFit.cover,
                progressIndicatorBuilder: (_, __, ___) {
                  return const SizedBox();
                },
              ),
              Padding(
                padding: const EdgeInsets.all(16.0),
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Row(
                      children: [
                        Text("Цены без скидок",
                          style: const TextStyle(
                            fontSize: 15,
                            fontWeight: FontWeight.w500,
                          ),
                        ),
                        Spacer(),
                        Text(
                          moneyFormatWithSymbol(logic.response!.priceBefore),
                          style: const TextStyle(
                            color: Colors.black,
                            fontSize: 15,
                            fontWeight: FontWeight.w500,
                            decoration: TextDecoration.lineThrough,
                          ),
                        ),
                      ],
                    ),
                    SizedBox(height: 8,),
                    Row(
                      children: [
                        Text("После цены:",
                          style: const TextStyle(
                            fontSize: 17,
                            fontWeight: FontWeight.w700,
                          ),
                        ),
                        Spacer(),
                        Text(
                          moneyFormatWithSymbol(logic.response!.priceAfter),
                          style: const TextStyle(
                            color: second,
                            fontSize: 17,
                            fontWeight: FontWeight.w700,
                          ),
                        ),
                      ],
                    ),
                    SizedBox(height: 8,),
                    Row(
                      children: [
                        Text("Скидка",
                          style: const TextStyle(
                            fontSize: 17,
                            fontWeight: FontWeight.w700,
                          ),
                        ),
                        Spacer(),
                        Text(
                          "${logic.response!.discountAmount} %",
                          style: const TextStyle(
                            color: first,
                            fontSize: 17,
                            fontWeight: FontWeight.w700,
                          ),
                        ),
                      ],
                    ),
                    SizedBox(height: 16,),
                    Row(
                      children: [
                        Text("Период скидки:",
                          style: const TextStyle(
                            fontSize: 17,
                            fontWeight: FontWeight.w700,
                          ),
                        ),
                        Spacer(),
                        Text("${logic.response?.fromTime?.split("T").first.replaceAll('-', '.')} - ${logic.response?.toTime?.split("T").first.replaceAll('-', '.')}",
                          style: const TextStyle(
                            fontSize: 17,
                            color: first,
                            fontWeight: FontWeight.w700,
                          ),
                        ),
                      ],
                    ),
                    SizedBox(height: 16,),
                    Row(
                      children: [
                        Text("Категория",
                          style: const TextStyle(
                            fontSize: 15,
                            fontWeight: FontWeight.w500,
                          ),
                        ),
                        Spacer(),
                        Text("${logic.response?.category?.name}",
                          style: const TextStyle(
                            fontSize: 15,
                            fontWeight: FontWeight.w500,
                          ),
                        ),
                      ],
                    ),
                    SizedBox(height: 16,),
                    Row(
                      children: [
                        Text("Производитель",
                          style: const TextStyle(
                            fontSize: 15,
                            fontWeight: FontWeight.w500,
                          ),
                        ),
                        Spacer(),
                        Text("${logic.response?.owner?.name}",
                          style: const TextStyle(
                            fontSize: 15,
                            fontWeight: FontWeight.w500,
                          ),
                        ),
                      ],
                    ),
                    SizedBox(height: 16,),
                    Text(logic.response!.description??'',
                      style: TextStyle(
                        fontSize: 15,
                        height: 1.3,
                      ),
                    )
                  ],
                ),
              ),
            ],
          ): SizedBox(),
        ),
        bottomNavigationBar: Padding(
          padding: EdgeInsets.all(16),
          child: MaterialButton(
            onPressed: (){
              launch(logic.response?.url??'');
            },
            shape: RoundedRectangleBorder(
              borderRadius: BorderRadius.circular(12),
            ),
            height: 56,
            color: second,
            child: Text("Покупка",
              style: TextStyle(
                color: Colors.white,
                fontWeight: FontWeight.w500,
                fontSize: 17
              ),
            ),
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
