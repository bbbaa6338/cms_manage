import os
import csv
import requests



cookies = {
    'x-token': 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiYjgyZjVhMmMtMzM2Yy00ZjZkLWI0MWMtNTM3MDI4NmFlNzUzIiwiSUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJOaWNrTmFtZSI6ImxhbGFsYWxhbGFsYWxhbGFsYSIsIkF1dGhvcml0eUlkIjo4ODgsIkJ1ZmZlclRpbWUiOjg2NDAwLCJpc3MiOiJxbVBsdXMiLCJhdWQiOlsiR1ZBIl0sImV4cCI6MTczMjUwOTQ1NCwibmJmIjoxNzMxOTA0NjU0fQ.AmyRmaOyzzci8b1veMDUG9YtvwAxDqG8PPihJCM19RM',
}

headers = {
    'Accept': 'application/json, text/plain, */*',
    'Accept-Language': 'zh-CN,zh;q=0.9,en;q=0.8',
    'Connection': 'keep-alive',
    'Content-Type': 'application/json',
    # 'Cookie': 'x-token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiYjgyZjVhMmMtMzM2Yy00ZjZkLWI0MWMtNTM3MDI4NmFlNzUzIiwiSUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJOaWNrTmFtZSI6ImxhbGFsYWxhbGFsYWxhbGFsYSIsIkF1dGhvcml0eUlkIjo4ODgsIkJ1ZmZlclRpbWUiOjg2NDAwLCJpc3MiOiJxbVBsdXMiLCJhdWQiOlsiR1ZBIl0sImV4cCI6MTczMjUwOTQ1NCwibmJmIjoxNzMxOTA0NjU0fQ.AmyRmaOyzzci8b1veMDUG9YtvwAxDqG8PPihJCM19RM',
    'Origin': 'http://38.181.29.193:33333',
    'Referer': 'http://38.181.29.193:33333/',
    'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36',
    'dnt': '1',
    'x-token': 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiYjgyZjVhMmMtMzM2Yy00ZjZkLWI0MWMtNTM3MDI4NmFlNzUzIiwiSUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJOaWNrTmFtZSI6ImxhbGFsYWxhbGFsYWxhbGFsYSIsIkF1dGhvcml0eUlkIjo4ODgsIkJ1ZmZlclRpbWUiOjg2NDAwLCJpc3MiOiJxbVBsdXMiLCJhdWQiOlsiR1ZBIl0sImV4cCI6MTczMjUwOTQ1NCwibmJmIjoxNzMxOTA0NjU0fQ.AmyRmaOyzzci8b1veMDUG9YtvwAxDqG8PPihJCM19RM',
    'x-user-id': '1',
}

sites = ["glzsbz.com", "wkstny.com", "zml1976.com", "xiongxincailiao.com", "zgznsq.com", "hgmrjtss.com", "skgdsb.com", "knwsfwx.com", "liqingf.com", "cqyfxx.com", "lyjlnk.com", "zzqjdc.com", "ylsqgj.com", "huazhongchaxun.com", "yfby888.com", "mayibanjia365.com", "tjsjgsbxg.com", "shy5188.com", "fullerence.com", "firedreamphoto.com", "shjd-edu.com", "slpaishuiban.com", "ynhrpzs.com", "zsb018.com", "zshongx.com", "shcdcc.com", "lcmygg.com", "hdsjxsb.com", "1cy37.com", "wsroujiamo.com", "jcks888.com", "wlguolv0038.com", "yachhf.com", "tjchangronggg.com", "bccsoy.com"]


# for site in sites:
#     json_data = {
#         'site': site,
#         'coreKeyword': site,
#         'updateArticleEveryDay': 5,
#         'updateStop': '',
#         'industry': '1',
#     }
#
#
#     response = requests.post(
#         'http://38.181.29.193:33333/api/information/createInformation',
#         cookies=cookies,
#         headers=headers,
#         json=json_data,
#         verify=False,
#     )
#
#     print(site, response.json())


# breakpoint()

start = 0

# 指定目录路径
directory_path = r'E:\\ciku\\二手车'

# 获取目录下所有的 CSV 文件
csv_files = [f for f in os.listdir(directory_path) if f.endswith('.csv')]

# 依次读取 CSV 文件
for csv_file in csv_files:
    file_path = os.path.join(directory_path, csv_file)
    print(f"Processing file: {csv_file}")
    subIndustry = file_path.split('\\')[-1].split('.')[0]
    subIndustry = subIndustry.replace('5118词库-二手车行业-', '')
    subIndustry = subIndustry.split('-')[0]

    with open(file_path, mode='r', encoding='gbk') as file:
        reader = csv.reader(file)
        # 跳过前两行（行索引从 0 开始，所以是 range(2)）
        for _ in range(2):
            next(reader, None)

        # 提取列 A（即每行的第一个元素）
        column_a = []

        for row in reader:
            if row:  # 确保行不为空
                r = row[0].strip()
                if len(r) > 2:
                    column_a.append(r)

        if len(column_a) > 4000:
            column_a = column_a[:4000]
        # 打印列 A 的数据
        # print("Column A Data:")
        # print(column_a)
        # print(column_a[0])
        # print(len(column_a))
        # print("-" * 40)  # 分隔线，便于阅读

        # print(sourceWord)
        # print(subIndustry)
        # break
        sourceWord = '\n'.join(column_a)

        site = sites[start]
        start += 1
        json_data = {
            'sourceWord': sourceWord,
            'industry': '1',
            'site': site,
            'subIndustry': subIndustry,
        }

        response = requests.post(
            'http://38.181.29.193:33333/api/industryKeyword/createIndustryKeyword',
            cookies=cookies,
            headers=headers,
            json=json_data,
            verify=False,
        )
        print(site, response.json())

