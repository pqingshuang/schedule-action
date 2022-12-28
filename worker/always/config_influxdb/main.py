# from config import Config, load_config as load, load_algo_X_wrapper_opt
# import tomconfig
# from functools import partial
import influxdb, datetime, time
# import win32com.client as win32
import argparse

import sys
# from chilkat_win_x64.chilkat import CkMailMan, CkEmail
from chilkat_linux.chilkat import CkMailMan, CkEmail

# ----------- Linux ----------- #
if __name__ == "__main__":
    # Instantiate the parser
    parser = argparse.ArgumentParser()

    parser.add_argument('--host', type=str)
    parser.add_argument('--port', type=int, default=8086)
    parser.add_argument('--database', type=str)
    parser.add_argument('--measurement', type=str)
    parser.add_argument('--time', type=str, default='1h')
    parser.add_argument('--email', type=str, default='marcus.law@neuroncloud.ai')

    args = parser.parse_args()

    #do your code with config
    #1. connect which influx
    #2. time 1h,
    #3. return True, None;
    # Flase,{}

    def send_mail(args, succ):
        mail_content = {'host': args.host, 'port': args.port, 'database': args.database, 'measurement': args.measurement}
        
        if succ:
            mail_head = 'Query successful at time', datetime.datetime.strftime(datetime.datetime.now(), '%d/%m/%Y %H:%M')
        else:
            mail_head = 'Query failed at time', datetime.datetime.strftime(datetime.datetime.now(), '%d/%m/%Y %H:%M')

            mailman = CkMailMan()

            smtpHost = mailman.mxLookup(args.email)
            if (mailman.get_LastMethodSuccess() != True):
                print(mailman.lastErrorText())
                sys.exit()
            mailman.put_SmtpHost(smtpHost)
            mail = CkEmail()
            mail.put_Subject(f"{mail_head}")
            mail.put_Body(f"{mail_content}")
            mail.put_From(args.email)
            mail.AddTo("", args.email)
            mailman.SendEmail(mail)
        # if (success != True):
        #     print(mailman.lastErrorText())
        # else:
        #     print("Sent!")

        return mail_content

    def CheckMeaExist(args):
        try:
            client = influxdb.DataFrameClient(f"{args.host}", args.port, database=f"{args.database}")
            influx_data = client.query(f"select * from {args.measurement} where time>now() - {args.time}")[f'{args.measurement}']
            if len(influx_data) > 0:
                mail_content = send_mail(args, True)
                return True, mail_content
            else:
                mail_content = send_mail(args, False)
                return False, mail_content

        except:
            mail_content = send_mail(args, False)
            return False, mail_content
            # print('Query failed at time', datetime.datetime.strftime(datetime.datetime.now(), '%d/%m/%Y %H:%M'), \
            #     f'with host: {args.host}, port: {args.port}, database: {args.database}, and measurement: {args.measurement}')

    CheckMeaExist(args)


# ----------- Windows ------------ #
# if __name__ == "__main__":
#     # Instantiate the parser
#     parser = argparse.ArgumentParser()

#     parser.add_argument('--host', type=str)
#     parser.add_argument('--port', type=int, default=8086)
#     parser.add_argument('--database', type=str)
#     parser.add_argument('--measurement', type=str)
#     parser.add_argument('--time', type=str, default='1h')
#     parser.add_argument('--sleeptime', type=int, default=3600)

#     args = parser.parse_args()

#     #do your code with config
#     #1. connect which influx
#     #2. time 1h,
#     #3. return True, None;
#     # Flase,{}

#     def CheckMeaExist(args):
#         try:
#             client = influxdb.DataFrameClient(f"{args.host}", args.port, database=f"{args.database}")
#             influx_data = client.query(f"select * from {args.measurement} where time>now() - {args.time}")[f'{args.measurement}']
#             if len(influx_data) > 0:
#                 print('Query successful at time', datetime.datetime.strftime(datetime.datetime.now(), '%d/%m/%Y %H:%M'))
#                 query_body = {'host': args.host, 'port': args.port, 'database': args.database, 'measurement': args.measurement}
#                 return True, query_body
#             else:
#                 outlook = win32.Dispatch('outlook.application')
#                 mail = outlook.CreateItem(0)
#                 mail.To = 'tom.lee@neuroncloud.ai'
#                 mail.Subject = f'Query Fail at time {datetime.datetime.strftime(datetime.datetime.now(), "%d/%m/%Y %H:%M")}'
#                 query_body = {'host': args.host, 'port': args.port, 'database': args.database, 'measurement': args.measurement}
#                 mail.Body = f"{query_body}"
#                 mail.Importance = 2
#                 mail.Send()
#                 return False, query_body
#                 # print('Query failed at time', datetime.datetime.strftime(datetime.datetime.now(), '%d/%m/%Y %H:%M'), \
#                 #     f'with host: {args.host}, port: {args.port}, database: {args.database}, and measurement: {args.measurement}')
#         except:
#             outlook = win32.Dispatch('outlook.application')
#             mail = outlook.CreateItem(0)
#             mail.To = 'tom.lee@neuroncloud.ai'
#             mail.Subject = f'Query Fail at time {datetime.datetime.strftime(datetime.datetime.now(), "%d/%m/%Y %H:%M")}'
#             query_body = {'host': args.host, 'port': args.port, 'database': args.database, 'measurement': args.measurement}
#             mail.Body = f"{query_body}"
#             mail.Importance = 2
#             mail.Send()
#             return False, query_body

#             # print('Query failed at time', datetime.datetime.strftime(datetime.datetime.now(), '%d/%m/%Y %H:%M'), \
#             #     f'with host: {args.host}, port: {args.port}, database: {args.database}, and measurement: {args.measurement}')

#     while True:
#         time.sleep(args.sleeptime)
#         CheckMeaExist(args)