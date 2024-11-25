#include "iofunc.h"

void image_set_kill(VipsImage *image) {
  vips_image_set_kill(image, 1);
}

int image_iskilled(VipsImage *image) {
  return vips_image_iskilled(image);
}
