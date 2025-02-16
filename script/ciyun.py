import jieba
from wordcloud import WordCloud
import matplotlib.pyplot as plt

# 读取文本文件
def read_text(file_path):
    with open(file_path, 'r', encoding='utf-8') as file:
        text = file.read()
    return text

# 中文分词
def chinese_word_segmentation(text):
    seg_list = jieba.cut(text)
    return " ".join(seg_list)

# 生成词云图
def generate_wordcloud(text):
    # 设置词云图的参数
    wc = WordCloud(
        font_path='simhei.ttf',  # 指定中文字体，确保中文能正常显示
        background_color='white',  # 设置背景颜色
        width=800,  # 设置词云图的宽度
        height=600,  # 设置词云图的高度
        max_words=200,  # 设置最大词数
        max_font_size=100  # 设置最大字体大小
    )
    # 生成词云
    wordcloud = wc.generate(text)

    # 显示词云图
    plt.figure(figsize=(8, 6))
    plt.imshow(wordcloud, interpolation='bilinear')
    plt.axis('off')  # 不显示坐标轴
    plt.show()

if __name__ == "__main__":
    # 文本文件的路径，需要替换为你自己的文件路径
    file_path = 'your_text_file.txt'
    # 读取文本
    text = read_text(file_path)
    # 中文分词
    segmented_text = chinese_word_segmentation(text)
    # 生成词云图
    generate_wordcloud("MYSQL ")