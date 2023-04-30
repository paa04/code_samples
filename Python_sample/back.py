import PIL
from PIL import Image


def get_code(s: str):  # Получить код символа
    s = ord(s)
    ans = [0] * 9
    for i in range(9):
        if (2 ** i) & s:
            ans[i] = 1
    return ans


def get_char(mas: list):  # Получить символ из пикселей
    ch = 0
    z = 1
    for i in mas:
        i &= 1
        ch += i * z
        z *= 2
    return chr(ch)


def write_pixel(mas: list, pix_int: int, image):  # записать символ
    for i in range(3):
        pix_num = pix_int + i
        x, y = image.size
        pix = list(image.getpixel((pix_num % x, pix_num // x)))
        pix = pix[:3]
        for j in range(3):
            pix[j] >>= 1
            pix[j] <<= 1
            pix[j] += mas[i * 3:i * 3 + 3][j]
        image.putpixel((pix_num % x, pix_num // x), tuple(pix))


def str_code(s, image):  # закодировать строку
    s += '\0'

    x, y = image.size

    if x * y < 3 * len(s):
        return None
    pix_num = 0
    for i in s:
        mas = get_code(i)
        write_pixel(mas, pix_num, image)
        pix_num += 3
    return image


def decode_pic(image):  # Раскодировать строку
    fl = False
    data = image.getdata()
    ans = ""
    for i in range(2, len(data), 3):
        ch = get_char(data[i - 2] + data[i - 1] + data[i])
        if ch == '\0':
            fl = True
            break
        ans += get_char(data[i - 2] + data[i - 1] + data[i])
    if fl:
        return ans
    else:
        return None
