import requests
from bs4 import BeautifulSoup

def crawl_baidu_images(keyword):
    # 构造百度图片搜索的 URL
    search_url = f'https://image.baidu.com/search/flip?tn=baiduimage&ie=utf-8&word={keyword}'

    # 发送 GET 请求获取页面内容
    response = requests.get(search_url)

    # 检查请求是否成功
    if response.status_code == 200:
        # 解析页面内容
        soup = BeautifulSoup(response.text, 'html.parser')

        # 找到图片链接的容器
        image_links = soup.find_all('img')

        # 遍历图片链接并打印
        for link in image_links:
            src = link.get('src')
            if src.startswith('http'):
                print(src)

    else:
        print("请求失败，状态码：", response.status_code)

# 示例用法
keyword = '花'
crawl_baidu_images(keyword)