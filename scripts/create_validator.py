import os
import sys
from getpass import getpass
import json
import subprocess

def generate_validator():
    print("Welcome to the Blit validator setup wizard!")
    print("This wizard will help you setup your validator.")
    print("Before we begin, please make sure you have the following:")

    print("Please enter your validator's moniker. This is the name that will be displayed on the explorer.")
    moniker = input("Moniker: ")

    print("Please enter your validator's identity like a UPort or Keybase signature. This is optional.")
    identity = input("Identity: ")

    print("Please enter your validator's website. This is optional.")
    website = input("Website: ")

    print("Please enter your validator's security contact email. This is optional.")
    security = input("Security: ")

    print("Please enter your validator's details. This is optional.")
    details = input("Details: ")

    print("Please enter your validator's commission rate. This is the percentage of rewards that will be charged to delegators.")
    commission_rate = input("Commission rate: [0.1] ") or "0.1"

    print("Please enter your validator's maximum commission rate. This is the maximum percentage of rewards that will be charged to delegators.")
    commission_max_rate = input("Commission max rate: [0.2] ") or "0.2"

    print("Please enter your validator's maximum commission change rate. This is the maximum % point change over the commission-rate. E.g. 1% to 2% is a 100% rate increase, but only 1 percentage point.")
    commission_max_change_rate = input("Commission max change rate: [0.01] ") or "0.01"

    print("Please enter your validator's minimum self delegation. This is the minimum amount of BLT that you will delegate to yourself.")
    min_self_delegation = input("Minimum self delegation: [1] ") or "1"

    print("Please enter your validator's amount. This is the amount of BLT that you will delegate to yourself.")
    amount = input("Amount: [1ublit] ") or "1ublit"

    # read pubkey from comet
    process = subprocess.Popen(["blitd", "comet", "show-validator"], stdout=subprocess.PIPE)
    output, error = process.communicate()
    pubkey = json.loads(output.decode("utf-8").strip())


    # write the validator.json file with the inputted values
    validator = {
        "pubkey": pubkey,
        "amount": amount,
        "moniker": moniker,
        "identity": identity,
        "website": website,
        "security": security,
        "details": details,
        "commission-rate": commission_rate,
        "commission-max-rate": commission_max_rate,
        "commission-max-change-rate": commission_max_change_rate,
        "min-self-delegation": min_self_delegation
    }

    return validator

if __name__ == "__main__":
    validator = generate_validator()
    print()
    print("Your validator data is:")
    print("====================================")
    print(json.dumps(validator, indent=2))
    print("====================================")
    save = input("Save this data to validator.json? [y/N]") or "n"
    if save in ["y", "Y"]:
        with open("validator.json", "w") as f:
            f.write(json.dumps(validator, indent=2))
        print("Saved to validator.json")
    
    print("Now you can run `blitd tx staking create-validator --help` to see how to create your validator.")




