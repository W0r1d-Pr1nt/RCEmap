import argparse
import urllib.parse

def construct_string(target_string):
    # 计算目标字符串的字节流
    target_bytes = bytes(target_string, 'latin1')
    # 计算取反后的字节流
    inverted_bytes = bytes(~byte & 0xff for byte in target_bytes)
    # 将取反后的字节流转换为字符串
    inverted_string = inverted_bytes.decode('latin1', errors='replace')
    # 对不可见字符进行 URL 编码
    encoded_string = ''.join(urllib.parse.quote(char) if char not in ' \t\n\r\f\v' else char for char in inverted_string)
    # 移除字符串中的 '%C2'
    encoded_string = encoded_string.replace('%C2', '')
    return encoded_string

def main():
    parser = argparse.ArgumentParser(description='字符串转换工具')
    parser.add_argument('target_string', help='要转换的字符串')
    args = parser.parse_args()

    constructed_string = construct_string(args.target_string)
    print(constructed_string)

if __name__ == "__main__":
    main()