import pandas as pd
import numpy as np
import influxdb
from tqdm import tqdm
import datetime
from itertools import chain
import argparse
from config_influxdb.chilkat_linux.chilkat import CkMailMan, CkEmail
argparser = argparse.ArgumentParser(
    description='Missing BMS_ID')

argparser.add_argument(
    '-t',
    '--time',
    default='1h',
    type=str,
    help='duration to get the data (in hour)')

argparser.add_argument(
    '-f',
    '--file',
    default='0',
    type=str,
    help='path to a default ID file')

def CreateIDFile():
    # check API_BIM_AM_Update_point.xlsx
    api_data = pd.ExcelFile('./API_BIM_AM_Update_point.xlsx')
    sheet_name = api_data.sheet_names  # see all sheet names
    id_names = []

    for name in sheet_name:
        api_data = pd.read_excel('./API_BIM_AM_Update_point.xlsx', sheet_name= name)[2:]
        id_names.append(list(api_data['Unnamed: 4']))

    id_names = list(chain.from_iterable(id_names))
    id_df = pd.DataFrame(id_names)
    id_df.rename(columns = {0:'id'}, inplace=True)
    # id_names.to_excel('./ID_List.xlsx', sheet_name='IDs', index=0)

    # id_df = id_names.set_index(id_names.columns[0])
    # print(id_df.head())

    return id_df

def GetIDNames(args):
    if args.file == '0':
        id_df = CreateIDFile()
    else:
        id_df = pd.ExcelFile(f'./{args.file}')
        id_df = id_df.parse(sheetname=id_df.sheet_names[0])
        # id_df = id_df.set_index(id_df.columns[0])

    print(id_df)
    return id_df

def CreateExcel(args, id_df):
    client = influxdb.DataFrameClient('192.168.100.213', 8086, database='NF')
    nf_data = client.query(f"Select * from Bms_db where time > now() - {args.time}")["Bms_db"]

    timenow = datetime.datetime.now()
    timenow = timenow.strftime('%d/%m/%Y %H:%M')
    id_df[timenow] = np.nan

    for i in tqdm(range(len(id_df))):
        for j in range(len(nf_data)-1, -1, -1):
            if nf_data['id'][j] == id_df['id'][i]:
                id_df[timenow][i] = nf_data['value'][j]
                break

    print('*' * 100)
    print(id_df.head())

    return id_df

def _main_(args):
    id_df = GetIDNames(args)
    curr_df = CreateExcel(args, id_df)
    filePath='./NF_IDData.xlsx'
    curr_df.to_excel(filePath, sheet_name='IDs', index=0)
    mailman = CkMailMan()

    smtpHost = mailman.mxLookup("hilbert.ng@neuroncloud.ai")
    if (mailman.get_LastMethodSuccess() != True):
        print(mailman.lastErrorText())
        sys.exit()
    mailman.put_SmtpHost(smtpHost)
    mail = CkEmail()
    mail.addFileAttachment(filePath)
    mail.put_Subject(f"NF Missing data daily checking")
    mail.put_Body(f"please check nf missing data attached ~ love you ")
    mail.put_From("hilbert.ng@neuroncloud.ai")
    mail.AddTo("", "hilbert.ng@neuroncloud.ai")
    mailman.SendEmail(mail)
    # print("good")


if __name__ == '__main__':
    args = argparser.parse_args()
    _main_(args)