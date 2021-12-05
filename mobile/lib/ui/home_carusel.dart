import 'package:cached_network_image/cached_network_image.dart';
import 'package:carousel_slider/carousel_slider.dart';
import 'package:chegirma/core/background_clipper.dart';
import 'package:chegirma/core/custom_pain.dart';
import 'package:flutter/material.dart';
import 'package:flutter_svg/svg.dart';
import 'package:get/get.dart';

class HomeCarousel extends StatelessWidget {
  const HomeCarousel({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return CarouselSlider(
        items: List.generate(
            3,
            (index) => InkWell(
                  splashColor: Colors.transparent,
                  onTap: () {},
                  child: Ink(
                    height: 208,
                    width: Get.width,
                    child: Stack(
                      children: [
                        CachedNetworkImage(
                          imageUrl:
                          data[index].image,
                          height: 208,
                          width: Get.width,
                          alignment: Alignment.bottomLeft,
                          errorWidget: (context, url, error) => Padding(
                            padding: const EdgeInsets.all(16),
                            child: SvgPicture.asset(
                              '',
                              color: Colors.white,
                            ),
                          ),
                          fit: BoxFit.cover,
                          progressIndicatorBuilder: (_, __, ___) {
                            return SizedBox();
                          },
                        ),
                        Positioned.fill(
                          child: CustomPaint(
                            painter: CustomPain(
                              data[index].color,
                            ),
                            child: ClipPath(
                              clipper: BackgroundClipper(),
                              child: Container(
                                color: data[index].seccolor,
                              ),
                            ),
                          ),
                        ),
                        Positioned(
                          right: 8,
                          bottom: 72,
                          child: Column(
                            crossAxisAlignment: CrossAxisAlignment.start,
                            mainAxisAlignment: MainAxisAlignment.center,
                            children: [
                              Container(
                                decoration: BoxDecoration(
                                  color: data[index].color,
                                  borderRadius: BorderRadius.circular(6),
                                ),
                                padding: const EdgeInsets.symmetric(vertical: 4, horizontal: 6),
                                margin: const EdgeInsets.only(bottom: 4),
                                child:  Text(
                                  data[index].title,
                                  style: const TextStyle(
                                    color: Colors.white,
                                  ),
                                ),
                              ),
                              Text(
                                data[index].subtitle,
                                style:  TextStyle(
                                  color: data[index].color,
                                ),
                              )
                            ],
                          ),
                        ),
                      ],
                    ),
                  ),
                )),
        options: CarouselOptions(
          height: 208,
          aspectRatio: 208 / Get.width,
          viewportFraction: 1,
          initialPage: 0,
          enableInfiniteScroll: true,
          reverse: false,
          scrollPhysics: const BouncingScrollPhysics(),
          autoPlay: true,
          autoPlayInterval: Duration(seconds: 6),
          autoPlayAnimationDuration: Duration(seconds: 6),
          autoPlayCurve: Curves.easeOut,
          onPageChanged: (_, index) {},
          scrollDirection: Axis.horizontal,
        ));
  }
}

class CaruselData{
  final String image;
  final String title;
  final Color color;
  final Color seccolor;
  final String subtitle;

  CaruselData({required this.image, required this.seccolor, required this.color, required this.title, required this.subtitle});
}

List<CaruselData> data = [
  CaruselData(
    subtitle: "попробуй и реши!",
    title: "Новый вкус",
    color: Colors.red,
    seccolor: Color(0x994B3032),
    image: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRn0MFAkOyrkyJ2Rbqe2rcQ67T_MwXZUc063Q&usqp=CAU",
  ),
  CaruselData(
    image: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTVfaEw7U9SFRKf-E8dYaD5yRWdJiC9SHH2pg&usqp=CAU",
    title: "При покупке две ",
    color: Colors.blue,
    seccolor: Color(0xAB01204A),
    subtitle: "Borjomi третий\nв подарок!",
  ),
  CaruselData(
    image: "https://www.depotwpf.ru/upload/media/u/2017/12/27/960_525_1%202.jpg",
    title: "При покупке две ",
    color: Colors.green,
    seccolor: Color(0x9900CC67),
    subtitle: "морожное третий\nв подарок!",
  )
];
