"""
Load the image based on URL

"""
import urllib.request
from urllib.error import HTTPError
import numpy as np
import csv


test_images = 'https://images-na.ssl-images-amazon.com/images/I/51ozrU27qyL._SL500_SR131,160_.jpg'

IMAGEURL_PATTERN = 'SL1000_SR640,640'


def load_image(url, image_asin=None, save_dir=''):
    """
    Based on the URL, get the images from the webpage.

    :param url:
    :return:
    """

    url_base = url[0:url.rfind('/')+1]
    image_id = url.split('/')[-1]
    image_base = image_id.split('.')[0]
    image_asin = image_asin if image_asin is not None else image_id
    new_image_url = url_base + image_base + '._{}_.jpg'.format(IMAGEURL_PATTERN)
    image_destination = save_dir + image_asin + '.jpg'
    try:
        url_response = urllib.request.urlretrieve(new_image_url, image_destination)
    except HTTPError:
        return None

    return url_response


def csv_reader(filepath, **kwargs):
    """

    Parameters
    ----------
    filepath
    kwargs : passed into csv reader

    Returns
    -------

    """
    # title = []
    entry = []
    with open(filepath, 'r') as f:
        reader = csv.reader(f, **kwargs)
        title = next(reader)
        for row in reader:
            entry.append(row)

    return title, entry


def process_images_with_rank():
    dataset_range = [71, 200]

    csv_path = '/Users/kcyu/Dropbox/go/src/github.com/hunterhug/AmazonBigSpider/result/test_images_1.csv'
    save_dir = '/Users/kcyu/dataset/shoes/images'

    title, entry = csv_reader(csv_path)
    print("title of CSV: {}".format(csv))
    for ind, e in enumerate(entry[dataset_range[0]: dataset_range[1]]):
        print("process {}: {}".format(ind, e))
        result = load_image(e[1], e[0], save_dir=save_dir)
        if result is None:
            continue



if __name__ == '__main__':
    process_images_with_rank()