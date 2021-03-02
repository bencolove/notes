# Shrik Image

1. Bitmap.compress only the data to be on disk
1. When BitmapFactory.decodeByteArray from compressed byte[], it will have same size as originally with poorer quality


```java
int len = bm.getByteCount();
ByteArrayOutputStream out = new ByteArrayOutputStream();
bm.compress(Bitmap.CompressFormat.JPEG, 80, out);

// output size shriks while
int compressedLen = out.toByteArray().length;
Bitmap compressedBm = BitmapFactory.decodeByteArray(out.toByteArray(), 0, compressedLen);
// stay the same in memory as Bitmap
int newLen = compressBm.getByteCount();
```

```
// 从选取相册的Activity中返回后
 
                    Uri imageUri = data.getData();
                    String[] filePathColumns = {MediaStore.Images.Media.DATA};
                    Cursor c = getContentResolver().query(imageUri, filePathColumns, null, null, null);
                    c.moveToFirst();
                    int columnIndex = c.getColumnIndex(filePathColumns[0]);
                    String imagePath = c.getString(columnIndex);
                    c.close();
                    
                    // 设置参数
                    BitmapFactory.Options options = new BitmapFactory.Options();
                    options.inJustDecodeBounds = true; // 只获取图片的大小信息，而不是将整张图片载入在内存中，避免内存溢出
                    BitmapFactory.decodeFile(imagePath, options);
                    int height = options.outHeight;
                    int width= options.outWidth;
                    int inSampleSize = 2; // 默认像素压缩比例，压缩为原图的1/2
                    int minLen = Math.min(height, width); // 原图的最小边长
                    if(minLen > 100) { // 如果原始图像的最小边长大于100dp（此处单位我认为是dp，而非px）
                        float ratio = (float)minLen / 100.0f; // 计算像素压缩比例
                        inSampleSize = (int)ratio;
                    }
                    options.inJustDecodeBounds = false; // 计算好压缩比例后，这次可以去加载原图了
                    options.inSampleSize = inSampleSize; // 设置为刚才计算的压缩比例
                    Bitmap bm = BitmapFactory.decodeFile(imagePath, options); // 解码文件
                    Log.w("TAG", "size: " + bm.getByteCount() + " width: " + bm.getWidth() + " heigth:" + bm.getHeight()); // 输出图像数据
                    imageView.setScaleType(ImageView.ScaleType.FIT_CENTER);
                    imageView.setImageBitmap(bm);
                    ```