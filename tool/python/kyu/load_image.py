"""
Load the image based on URL

"""
import urllib.request
import numpy as np

test_images = 'https://images-na.ssl-images-amazon.com/images/I/51ozrU27qyL._SL500_SR131,160_.jpg'

IMAGEURL_PATTERN = 'SL1000_SR640,640'


def load_image(url):
    """
    Based on the URL, get the images from the webpage.

    :param url:
    :return:
    """
    # TODO find a better encoding
    url_base = url.split('/')[0:-1]
    image_id = url.split('/')[-1]

    url_response = urllib.request.urlretrieve(url, )
    img_array = np.array()


